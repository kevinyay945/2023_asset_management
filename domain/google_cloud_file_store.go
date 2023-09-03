package domain

type GoogleCloudFileStore struct{}

func NewGoogleCloudFileStore() CloudFileStorer {
	return &GoogleCloudFileStore{}
}

func (c *GoogleCloudFileStore) GetPublicLink(file CloudFile) (link string, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *GoogleCloudFileStore) UploadFile(name string, mimeType string, data []byte, location CloudFileLocation) (file CloudFile, err error) {
	err = nil
	return
}
