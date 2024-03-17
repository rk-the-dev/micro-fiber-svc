package securityhelper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
	"os"
)

type Securityhelper struct {
}

func Getsecurityhelper() *Securityhelper {
	return &Securityhelper{}
}

// getEncryptionKey retrieves the AES encryption key from an environment variable.
func getEncryptionKey() []byte {
	key := os.Getenv("AES_ENCRYPTION_KEY")
	if len(key) == 0 {
		log.Fatal("AES_ENCRYPTION_KEY environment variable is not set or is invalid.")
	}
	return []byte(key)
}

// encrypt encrypts plain text string using AES.
func (s *Securityhelper) Encrypt(plainText string) (string, error) {
	block, err := aes.NewCipher(getEncryptionKey())
	if err != nil {
		return "", err
	}
	plainTextBytes := []byte(plainText)
	cipherText := make([]byte, aes.BlockSize+len(plainTextBytes))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainTextBytes)

	return hex.EncodeToString(cipherText), nil
}

// decrypt decrypts cipher text string using AES.
func (s *Securityhelper) Decrypt(cipherText string) (string, error) {
	block, err := aes.NewCipher(getEncryptionKey())
	if err != nil {
		return "", err
	}
	cipherTextBytes, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	if len(cipherTextBytes) < aes.BlockSize {
		return "", err
	}
	iv := cipherTextBytes[:aes.BlockSize]
	cipherTextBytes = cipherTextBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherTextBytes, cipherTextBytes)

	return string(cipherTextBytes), nil
}
