package filestorage

import "io"

type directoryStorage struct {
	basePath string
}

func NewDirectoryStorage(basePath string) FileStorage {
	return &directoryStorage{basePath: basePath}
}

func (fs *directoryStorage) NewTransaction() (io.WriteCloser, error) {
	return nil, nil
}
