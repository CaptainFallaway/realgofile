package hasher

import (
	"crypto/rand"
)

type HasherWithSalt interface {
	Hash(plain string, salt []byte) ([]byte, error)
	Compare(plain string, salt, hashed []byte) bool
}

type Hasher interface {
	Hash(plain string) ([]byte, error)
	Compare(plain string, hashed []byte) bool
}

func GenerateSalt(saltSize uint8) ([]byte, error) {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		return nil, err
	}

	return salt, nil
}
