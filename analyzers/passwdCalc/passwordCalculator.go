package passwdCalc

import (
	"MobileForensicsToolkit/utils"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// PasswordCalculator handles password generation for various messaging applications
type PasswordCalculator struct {
	ctx context.Context
}

// NewPasswordCalculator creates a new PasswordCalculator instance
func NewPasswordCalculator() *PasswordCalculator {
	return &PasswordCalculator{}
}

// InitCtx initializes the context for the calculator
func (p *PasswordCalculator) InitCtx(ctx context.Context) {
	p.ctx = ctx
}

// startup initializes the calculator at startup
func (p *PasswordCalculator) startup(ctx context.Context) {
	p.ctx = ctx
}

// unpad removes PKCS7 padding from decrypted data
func unpad(ciphertext []byte) ([]byte, error) {
	length := len(ciphertext)
	if length%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext block size is invalid")
	}
	padding := ciphertext[length-1]
	if int(padding) > length {
		return nil, errors.New("padding is invalid")
	}
	return ciphertext[:length-int(padding)], nil
}

// decryptAES decrypts AES-encrypted base64 token using CBC mode
func decryptAES(token string, key, iv []byte) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	plaintext, err := unpad(ciphertext)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// CalWechat calculates WeChat database decryption key
func (p *PasswordCalculator) CalWechat(uin, imei string) string {
	/*
		@param uin: User ID
		@param imei: Phone IMEI, defaults to "1234567890ABCDEF" if empty
		@return: Decryption key for Android WeChat EncMicroMsg.db
	*/
	if imei == "" {
		imei = "1234567890ABCDEF"
	}
	pwd := utils.MD5HashHex(imei + uin)[:7]
	return pwd
}

// CalWechatIndex calculates WeChat index database decryption key
func (p *PasswordCalculator) CalWechatIndex(uin, wxid, imei string) string {
	/*
		@param uin: User ID
		@param imei: Phone IMEI, defaults to "1234567890ABCDEF" if empty
		@param wxid: WeChat internal user ID (wxid)
		@return: Decryption key for Android WeChat index databases
	*/
	if imei == "" {
		imei = "1234567890ABCDEF"
	}
	atoi, err := strconv.Atoi(uin)
	if err != nil {
		return ""
	}
	if atoi < 0 {
		atoi += 4294967296
	}
	uin = strconv.Itoa(atoi)
	return utils.MD5HashHex(uin + imei + wxid)[:7]
}

// CalWildFire calculates WildFire IM application database decryption key
func (p *PasswordCalculator) CalWildFire(token string) []string {
	/*
		@param token: WildFire IM application user token, found in shared_prefs/config.xml
		@return: Array containing decryption key and instructions
	*/
	// New version key
	key, _ := hex.DecodeString("001122334455667778797A7B7C7D7E7F")
	iv := key
	// Old version key
	key2, _ := hex.DecodeString("7F7E7D7C7B7A79787766554433221100")
	iv2 := key2
	
	pwd, err := decryptAES(token, key, iv)
	flag := "Use SQLCipher4 for decryption"
	if err != nil {
		pwd, err = decryptAES(token, key2, iv2)
		if err != nil {
			return []string{"Decryption failed", ""}
		}
		flag = "Use SQLCipher3 for decryption"
	}
	parts := strings.Split(pwd, "|")
	if len(parts) > 0 {
		pwd = parts[len(parts)-1]
	}
	return []string{pwd, flag}
}

// CalMostone calculates Mostone (默往) messaging app database decryption key
func (p *PasswordCalculator) CalMostone(uid string) []string {
	/*
		@param uid: Mostone user UID, found in shared_prefs/im.xml userId value
		@return: Array containing decryption key for msg.db and instructions
	*/
	hash := utils.MD5HashHex(uid)
	ret := strings.ToUpper(hash[:6])
	return []string{ret, "Use SQLCipher3 for decryption"}
}

// CalTiktok calculates TikTok (抖音) database decryption key
func (p *PasswordCalculator) CalTiktok(uid string) []string {
	/*
		@param uid: TikTok user UID, usually found in database filename
		@return: Array containing decryption key and instructions
	*/
	en := fmt.Sprintf("byte%simwcdb%sdance", uid, uid)
	return []string{en, "Use WCDB for decryption"}
}

// CalXiaoHongShu calculates XiaoHongShu (小红书) database decryption key
func (p *PasswordCalculator) CalXiaoHongShu() []string {
	/*
		@return: Array containing decryption key and instructions for XiaoHongShu
	*/
	return []string{"xhsdev", "Use SQLCipher3 for decryption"}
}

// CalMosGram calculates MosGram/Bubble (泡泡) database decryption key
func (p *PasswordCalculator) CalMosGram(custId string) []string {
	/*
		@param custId: Customer ID from account_config.xml in sp directory
		@return: Array containing MD5 hash of custId and instructions
	*/
	pwd := utils.MD5HashHex(custId)
	return []string{pwd, "Use SQLCipher4 for decryption"}
}

// CalWuKongIM calculates WuKong IM series database decryption key
func (p *PasswordCalculator) CalWuKongIM(userId string) []string {
	/*
		@param userId: User ID, which is also the database password
		@return: Array containing user ID as password and instructions
	*/
	return []string{userId, "Use SQLCipher4 for decryption. Database name: wk_" + userId + ".db"}
}
