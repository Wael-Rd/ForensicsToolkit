package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// HashType defines the hash algorithm type
type HashType int

const (
	MD5 HashType = iota
	SHA1
	SHA256
)

// Hash calculates the hash value of input data
// Input can be string or byte array
// Returns byte array
func Hash(data interface{}, hashType HashType) []byte {
	var input []byte

	switch v := data.(type) {
	case string:
		input = []byte(v)
	case []byte:
		input = v
	default:
		return nil
	}

	switch hashType {
	case MD5:
		h := md5.Sum(input)
		return h[:]
	case SHA1:
		h := sha1.Sum(input)
		return h[:]
	case SHA256:
		h := sha256.Sum256(input)
		return h[:]
	default:
		return nil
	}
}

// HashHex calculates the hash value of input data and returns hex string
func HashHex(data interface{}, hashType HashType) string {
	hashBytes := Hash(data, hashType)
	return hex.EncodeToString(hashBytes)
}

// Shortcut functions for specific algorithms

// MD5Hash calculates MD5 hash value
func MD5Hash(data interface{}) []byte {
	return Hash(data, MD5)
}

// MD5HashHex calculates MD5 hash value and returns hex string
func MD5HashHex(data interface{}) string {
	return HashHex(data, MD5)
}

// SHA1Hash calculates SHA1 hash value
func SHA1Hash(data interface{}) []byte {
	return Hash(data, SHA1)
}

// SHA1HashHex calculates SHA1 hash value and returns hex string
func SHA1HashHex(data interface{}) string {
	return HashHex(data, SHA1)
}

// SHA256Hash calculates SHA256 hash value
func SHA256Hash(data interface{}) []byte {
	return Hash(data, SHA256)
}

// SHA256HashHex calculates SHA256 hash value and returns hex string
func SHA256HashHex(data interface{}) string {
	return HashHex(data, SHA256)
}
