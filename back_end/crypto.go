package back_end

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(str string) (hash string, err error) {
	var b []byte
	b, err = bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	hash = string(b)
	return
}

func Compare(hash string, str string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	if err != nil {
		return false
	}
	return true
}
