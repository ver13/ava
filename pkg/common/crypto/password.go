package crypto

import (
	"fmt"
	
	"golang.org/x/crypto/bcrypt"
	
	errCryptoAVA "github.com/ver13/ava/pkg/common/crypto/error"
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type Password string

func NewPassword(password string) PasswordI {
	return Password(password)
}

func (p Password) Hash() (string, *errorAVA.Error) {
	var hash []byte

	if len(p) > 0 {
		var err error
		hash, err = bcrypt.GenerateFromPassword([]byte(p), 10)
		if err != nil {
			return "", errCryptoAVA.BcryptHash(err, "")
		}
	}
	hashPassword := string(hash)

	return hashPassword, nil
}

func (p Password) ComparePassword(newPassword string) *errorAVA.Error {
	if len(p) == 0 || len(newPassword) == 0 {
		return errCryptoAVA.PasswordIsEmpty(nil, fmt.Sprintf("Password: %s - NewPassword: %s.", p, newPassword))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(p), []byte(newPassword)); err != nil {
		return errCryptoAVA.MismatchedHashAndPassword(err, "Hashed password is not the hash of the given password.")
	}

	return nil
}
