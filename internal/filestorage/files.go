package filestorage

import "io"

type directoryStorage struct {
	basePath string
}

func NewDirectoryStorage(basePath string) FileStorage {
	return &directoryStorage{basePath: basePath}
}

func (ds *directoryStorage) Delete(uid string) error {
	return nil
}

func (ds *directoryStorage) NewWriteTransaction(uid string) (io.WriteCloser, error) {
	return nil, nil
}

func (ds *directoryStorage) NewReadTransaction(uid string) (io.ReadCloser, error) {
	return nil, nil
}
