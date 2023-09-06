package application

import "2023_asset_management/domain"

//go:generate mockgen -destination=file_store.mock.go -package=application -self_package=2023_asset_management/application . FileStorer
type FileStorer interface {
	UploadAsset(fileName string, data []byte, location domain.CloudFileLocation, mimeType string) (file domain.CloudFile, err error)
	GetPreviewLink(file domain.CloudFile) (link string, err error)
	GetPreviewLinkByLocationAndFileName(location domain.CloudFileLocation, fileName string) (link string, err error)
}

type FileStore struct {
	CloudFileStore      domain.CloudFileStorer
	PreviewLinkGenerate domain.PreviewLinkGenerator
}

func NewFileStore(cloudFileStore domain.CloudFileStorer, previewLinkGenerate domain.PreviewLinkGenerator) FileStorer {
	return &FileStore{CloudFileStore: cloudFileStore, PreviewLinkGenerate: previewLinkGenerate}
}

func (f *FileStore) GetPreviewLinkByLocationAndFileName(location domain.CloudFileLocation, fileName string) (link string, err error) {
	file, err := f.CloudFileStore.GetCloudFileByNameAndLocation(fileName, location)
	if err != nil {
		return
	}
	link, err = f.CloudFileStore.GetPublicLink(file)
	if err != nil {
		return
	}
	return
}

func (f *FileStore) UploadAsset(fileName string, data []byte, location domain.CloudFileLocation, mimeType string) (file domain.CloudFile, err error) {
	file, err = f.CloudFileStore.UploadFile(fileName, mimeType, data, location)
	return
}

func (f *FileStore) GetPreviewLink(file domain.CloudFile) (link string, err error) {
	link = f.PreviewLinkGenerate.GetLinkByCloudFile(file)
	return
}
