package hashing

import (
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
)

type argonBcryptHasher struct {
	bcryptCost   int
	argonMemory  uint32
	argonThreads uint8
	argonTime    uint32
}

func NewDefaultArgonBcryptHasher() HasherWithSalt {
	return &argonBcryptHasher{
		bcryptCost:   12,
		argonMemory:  64 * 1024,
		argonThreads: 4,
		argonTime:    1,
	}
}

func NewArgonBcryptHasher(bcryptCost int, argonMemory uint32, argonThreads uint8, argonTime uint32) HasherWithSalt {
	return &argonBcryptHasher{bcryptCost, argonMemory, argonThreads, argonTime}
}

func (a *argonBcryptHasher) argonHash(plain string, salt []byte, keyLen uint32) []byte {
	return argon2.IDKey([]byte(plain), salt, a.argonTime, a.argonMemory, a.argonThreads, keyLen)
}

func (a *argonBcryptHasher) Hash(plain string, salt []byte) ([]byte, error) {
	hashedKey := a.argonHash(plain, salt, 72)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(hashedKey), a.bcryptCost)
	if err != nil {
		return nil, err
	}

	return hashedPassword, nil
}

func (a *argonBcryptHasher) Compare(plainKey string, salt, hashed []byte) bool {
	argonHash := a.argonHash(plainKey, salt, 72)
	return bcrypt.CompareHashAndPassword(hashed, argonHash) == nil
}
