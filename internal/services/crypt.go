package services

type CryptService interface {
	Encrypt(key string, salt, data []byte) ([]byte, error)
	Decrypt(key string, salt []byte) ([]byte, error)
}
