package application

import "2023_asset_management/domain"

//go:generate mockgen -destination=file_store.mock.go -package=application -self_package=2023_asset_management/application . FileStorer
type FileStorer interface {
	UploadAsset(fileName string, data []byte, location domain.CloudFileLocation) (file domain.CloudFile, err error)
	GetPreviewLink(file domain.CloudFile) (link string, err error)
}

type FileStore struct {
	CloudFileStore domain.CloudFileStorer
}

func NewFileStore() FileStorer {
	return &FileStore{}
}

func (f *FileStore) UploadAsset(fileName string, data []byte, location domain.CloudFileLocation) (file domain.CloudFile, err error) {
	file, err = f.CloudFileStore.UploadFile(fileName, "image/png", data, location)
	return
}

func (f *FileStore) GetPreviewLink(file domain.CloudFile) (link string, err error) {
	link, err = f.CloudFileStore.GetPublicLink(file)
	return
}
