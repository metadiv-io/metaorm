package metaorm

import (
	"os"

	"github.com/metadiv-io/aes"
)

// Encrypt encrypts the given text using the METAORM_ENCRYPT_KEY environment variable.
func Encrypt(text string) []byte {
	return aes.Encrypt.StringToBytes(text, os.Getenv("METAORM_ENCRYPT_KEY"))
}

// Decrypt decrypts the given cipher text using the METAORM_ENCRYPT_KEY environment variable.
func Decrypt(cipherText []byte) string {
	return aes.Decrypt.BytesToString(cipherText, os.Getenv("METAORM_ENCRYPT_KEY"))
}
