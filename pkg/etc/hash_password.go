package etc

import (
	"crypto/rand"
	"io"

	"golang.org/x/crypto/bcrypt"
)

var (
	// table for code generator
	table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
	// teng bomasa error nilga teng bomaydi va false qaytaradi
	// teng bosa error nilga teng boladi va true qaytaradi
}

func GenerateCode(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)

}
