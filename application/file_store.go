package application

import "2023_asset_management/domain"

//go:generate mockgen -destination=file_store_mock.go -package=application -self_package=2023_asset_management/application . FileStorer
type FileStorer interface {
	UploadAsset(filename string, i []byte, s string) (file domain.CloudFile, err error)
	GetPreviewLink(asset domain.CloudFile) (link string, err error)
}

type FileStore struct {
}

func NewFileStore() FileStorer {
	return &FileStore{}
}

func (f *FileStore) UploadAsset(filename string, i []byte, s string) (file domain.CloudFile, err error) {
	//TODO implement me
	panic("implement me")
}

func (f *FileStore) GetPreviewLink(asset domain.CloudFile) (link string, err error) {
	//TODO implement me
	panic("implement me")
}
