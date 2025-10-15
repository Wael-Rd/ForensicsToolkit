package database

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/WXjzcccc/go-sqlcipher"
)

// Database decryption pragma constants for different encryption types
const (
	EnMicroMsgDecryptPragma = "_pragma_cipher_compatibility=1"
	FTSIndexDecryptPragma   = "_pragma_kdf_iter=64000&_pragma_cipher_kdf_algorithm=PBKDF2_HMAC_SHA1&_pragma_cipher_hmac_algorithm=HMAC_SHA1"
	SQLCipher4Pragma        = ""
	SQLCipher3Pragma        = "_pragma_cipher_page_size=1024&_pragma_kdf_iter=64000&_pragma_cipher_kdf_algorithm=PBKDF2_HMAC_SHA1&_pragma_cipher_hmac_algorithm=HMAC_SHA1"
	WcdbPragma              = "_pragma_cipher_page_size=4096&_pragma_kdf_iter=64000&_pragma_cipher_kdf_algorithm=PBKDF2_HMAC_SHA1&_pragma_cipher_hmac_algorithm=HMAC_SHA1"
	NtqqPragma              = "_pragma_cipher_page_size=4096&_pragma_kdf_iter=4000&_pragma_cipher_kdf_algorithm=PBKDF2_HMAC_SHA512"
)

// Database type constants
const (
	FTSIndexDB   = 0
	SQLCipher4DB = 1
	SQLCipher3DB = 2
	WCdbDB       = 3
)

// DecryptDatabase handles database decryption operations
type DecryptDatabase struct {
	ctx context.Context
}

// NewDecryptDatabase creates a new DecryptDatabase instance
func NewDecryptDatabase() *DecryptDatabase {
	return &DecryptDatabase{}
}

// DecryptResult represents the result of a database decryption operation
type DecryptResult struct {
	SavePath string `json:"save_path"` // Path to the decrypted database file
	Wxid     string `json:"wxid"`      // WeChat ID (for WeChat databases)
	Err      string `json:"err"`       // Error message if decryption failed
}

// InitCtx initializes the context for the analyzer
func (d *DecryptDatabase) InitCtx(ctx context.Context) {
	d.ctx = ctx
}

// moveElementToFirst moves a specific element to the first position in a slice
func moveElementToFirst(slice []string, element string) []string {
	idx := 0
	for i, v := range slice {
		if v == element {
			idx = i
			break
		}
	}
	// Create a new slice with the same length as the original
	newSlice := make([]string, len(slice))

	// Place the target element at the first position
	newSlice[0] = slice[idx]

	copy(newSlice[1:idx+1], slice[0:idx])
	copy(newSlice[1+idx:], slice[idx+1:])

	return newSlice
}

// decryptSql generates SQL command for database decryption export
func decryptSql(dbPath string) (string, string) {
	dbName := "a" + strings.Replace(filepath.Base(dbPath), ".", "", -1) // Variables cannot start with numbers
	savePath := dbPath + "_dec.db"
	return fmt.Sprintf("ATTACH DATABASE '%s' AS '%s_dec' KEY '';SELECT sqlcipher_export('%s_dec');DETACH DATABASE '%s_dec';", savePath, dbName, dbName, dbName), savePath
}

// DecryptEnMicroMsg decrypts WeChat EnMicroMsg database
func (d *DecryptDatabase) DecryptEnMicroMsg(dbPath, password string) *DecryptResult {
	/*
		@param dbPath	: Database file path
		@param password	: Decryption password
		@return		 	: Decrypted database path, wxid, error
	*/
	fileInfo, err := os.Stat(dbPath)
	if err != nil {
		return &DecryptResult{"", "", err.Error()}
	}
	if fileInfo.IsDir() {
		return &DecryptResult{"", "", fmt.Sprintf("%s is not a file", dbPath)}
	}
	pwd := url.QueryEscape(password)
	dbName := fmt.Sprintf("%s?_key=%s&%s", dbPath, pwd, EnMicroMsgDecryptPragma)
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return &DecryptResult{"", "", fmt.Sprintf("Failed to open %s: %v", dbPath, err)}
	}
	defer db.Close()
	
	// Get WeChat ID
	selectSql := "select value from userinfo where id = 2;"
	cur, err := db.Prepare(selectSql)
	if err != nil {
		return &DecryptResult{"", "", fmt.Sprintf("Failed to prepare SQL [%s]: %v", selectSql, err)}
	}
	var wxid string
	err = cur.QueryRow().Scan(&wxid)
	if err != nil {
		return &DecryptResult{"", "", fmt.Sprintf("Failed to read wxid: %v", err)}
	}
	decCmd, savePath := decryptSql(dbPath)
	_, err = db.Exec(decCmd)
	if err != nil {
		return &DecryptResult{"", "", fmt.Sprintf("Failed to decrypt database: %v", err)}
	}
	return &DecryptResult{savePath, wxid, ""}
}

// decryptNormal decrypts standard encrypted databases
func (d *DecryptDatabase) decryptNormal(dbPath, password string, dbType int) *DecryptResult {
	/*
		@param dbPath	: Database file path
		@param password	: Decryption password
		@param dbType	: Database encryption type
		@return		 	: Decrypted database path, wxid (always empty for non-WeChat), error
	*/
	fileInfo, err := os.Stat(dbPath)
	if err != nil {
		return &DecryptResult{"", "", err.Error()}
	}
	if fileInfo.IsDir() {
		return &DecryptResult{"", "", fmt.Sprintf("%s is not a file", dbPath)}
	}
	var dbName string
	pwd := url.QueryEscape(password)
	switch dbType {
	case FTSIndexDB:
		dbName = fmt.Sprintf("%s?_key=%s&%s", dbPath, pwd, FTSIndexDecryptPragma)
	case SQLCipher4DB:
		dbName = fmt.Sprintf("%s?_key=%s", dbPath, pwd)
	case SQLCipher3DB:
		dbName = fmt.Sprintf("%s?_key=%s&%s", dbPath, pwd, SQLCipher3Pragma)
	case WCdbDB:
		dbName = fmt.Sprintf("%s?_key=%s&%s", dbPath, pwd, WcdbPragma)
	default:
		return &DecryptResult{"", "", fmt.Sprintf("Unsupported database type: %d", dbType)}
	}
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return &DecryptResult{"", "", fmt.Sprintf("Failed to open %s: %v", dbPath, err)}
	}
	defer db.Close()
	decCmd, savePath := decryptSql(dbPath)
	_, err = db.Exec(decCmd)
	if err != nil {
		return &DecryptResult{"", "", fmt.Sprintf("Failed to decrypt database: %v", err)}
	}
	return &DecryptResult{savePath, "", ""}
}

// DecryptFTSIndex decrypts FTS index databases
func (d *DecryptDatabase) DecryptFTSIndex(dbPath, password string) *DecryptResult {
	return d.decryptNormal(dbPath, password, FTSIndexDB)
}

// DecryptSQLCipher4 decrypts SQLCipher 4 databases
func (d *DecryptDatabase) DecryptSQLCipher4(dbPath, password string) *DecryptResult {
	return d.decryptNormal(dbPath, password, SQLCipher4DB)
}

// DecryptSQLCipher3 decrypts SQLCipher 3 databases  
func (d *DecryptDatabase) DecryptSQLCipher3(dbPath, password string) *DecryptResult {
	return d.decryptNormal(dbPath, password, SQLCipher3DB)
}

// DecryptWCdb decrypts WCDB databases
func (d *DecryptDatabase) DecryptWCdb(dbPath, password string) *DecryptResult {
	return d.decryptNormal(dbPath, password, WCdbDB)
}

// Additional methods would continue here following the same pattern...
// The file continues with more encryption/decryption methods for various database types
