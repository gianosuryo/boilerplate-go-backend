package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"os"
	"sarana-dafa-ai-service/storage/env"
)

/**
* Use: NewCTR
**/
func DecryptSalesToken(encryptedData string) (string, error) {
	key := []byte(os.Getenv(env.SALES_SECRET))
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	plaintext := make([]byte, len(ciphertext))

	ctr := cipher.NewCTR(block, iv)
	ctr.XORKeyStream(plaintext, ciphertext)
	return string(plaintext), nil
}

/**
* Use: NewCFBEncrypter
**/
func DecryptToken(token string) (string, error) {
	key := []byte(os.Getenv(env.ENCRYPTION_SECRET))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)

	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}
func EncryptToken(plainText string) (string, error) {
	key := []byte(os.Getenv(env.ENCRYPTION_SECRET))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(plainText))

	encoded := base64.StdEncoding.EncodeToString(cipherText)
	return encoded, nil
}
