package hasher

import (
	"bytes"
	"crypto/sha1"
)

type shaHasher struct{}

func NewShaHasher() Hasher {
	return &shaHasher{}
}

func (h *shaHasher) Hash(plain string) ([]byte, error) {
	hash := sha1.New()
	return hash.Sum([]byte(plain)), nil
}

func (h *shaHasher) Compare(plain string, hashed []byte) bool {
	hash := sha1.New()
	hashedPlain := hash.Sum([]byte(plain))
	return bytes.Equal(hashedPlain, hashed)
}
