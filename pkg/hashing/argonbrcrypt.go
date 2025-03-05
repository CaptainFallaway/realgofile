package hashing

import (
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost   = 12
	argonMemory  = 64 * 1024
	argonThreads = 4
	argonTime    = 1
)

type argonBcryptHasher struct{}

func NewArgonBcryptHasher() HasherWithSalt {
	return &argonBcryptHasher{}
}

func argonHash(plain string, salt []byte, keyLen uint32) []byte {
	return argon2.IDKey([]byte(plain), salt, argonTime, argonMemory, argonThreads, keyLen)
}

func (a *argonBcryptHasher) Hash(plain string, salt []byte) ([]byte, error) {
	hashedKey := argonHash(plain, salt, 72)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(hashedKey), bcryptCost)
	if err != nil {
		return nil, err
	}

	return hashedPassword, nil
}

func (a *argonBcryptHasher) Compare(plainKey string, salt, hashed []byte) bool {
	argonHash := argonHash(plainKey, salt, 72)
	return bcrypt.CompareHashAndPassword(hashed, argonHash) == nil
}
