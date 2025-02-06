package back_end

import (
	"golang.org/x/crypto/bcrypt"
)

func encryptPassword(password string) (hash string, err error) {
	var b []byte
	b, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hash = string(b)
	return
}
