package metaorm

import (
	"os"

	"github.com/metadiv-io/aes"
)

// Encrypt encrypts the given text using the given encrypt key.
func Encrypt(text string, encryptKey string) []byte {
	return aes.Encrypt.StringToBytes(text, encryptKey)
}

// DefaultEncrypt encrypts the given text using the METAORM_ENCRYPT_KEY environment variable.
func DefaultEncrypt(text string) []byte {
	return Encrypt(text, os.Getenv("METAORM_ENCRYPT_KEY"))
}

// Decrypt decrypts the given cipher text using the given decrypt key.
func Decrypt(cipherText []byte, decryptKey string) string {
	return aes.Decrypt.BytesToString(cipherText, decryptKey)
}

// DefaultDecrypt decrypts the given cipher text using the METAORM_ENCRYPT_KEY environment variable.
func DefaultDecrypt(cipherText []byte) string {
	return Decrypt(cipherText, os.Getenv("METAORM_ENCRYPT_KEY"))
}
