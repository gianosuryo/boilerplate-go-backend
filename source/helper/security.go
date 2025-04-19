package helper

import (
	"crypto/sha1"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) (string, error) {
	// md5Hashed := md5.Sum([]byte(os.Getenv(env.PASSWORD_SECRET) + password))
	// md5HashedString := hex.EncodeToString(md5Hashed[:])
	// bytes, err := bcrypt.GenerateFromPassword([]byte(md5HashedString), bcrypt.DefaultCost)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
func PasswordVerify(hashedPassword string, password string) bool {
	// md5Hashed := md5.Sum([]byte(os.Getenv(env.PASSWORD_SECRET) + password))
	// md5HashedString := hex.EncodeToString(md5Hashed[:])
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func GenerateStreamToken(generateId string) string {
	textStream := fmt.Sprintf("generateId:%s&Unix:%d", generateId, time.Now().UnixNano())
	sha := sha1.New()
	sha.Write([]byte(textStream))

	return fmt.Sprintf("%x", sha.Sum(nil))
}
