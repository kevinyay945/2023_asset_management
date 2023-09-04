package domain

import (
	"2023_asset_management/infrastructure/googledrive"
	"google.golang.org/api/drive/v3"
)

type GoogleCloudFileStore struct {
	googleDriver googledrive.GoogleDriver
}

func NewGoogleCloudFileStore(googleDriver googledrive.GoogleDriver) CloudFileStorer {
	return &GoogleCloudFileStore{googleDriver: googleDriver}
}

func (c *GoogleCloudFileStore) GetCloudFileByName(name string) (file CloudFile, err error) {
	file.Name = name
	return
}

func (c *GoogleCloudFileStore) GetPublicLink(file CloudFile, location CloudFileLocation) (link string, err error) {
	link, err = c.googleDriver.GetFilePublicLink(location.GoogleDriveUploadLocation(), file.Name)
	return
}

func (c *GoogleCloudFileStore) UploadFile(name string, mimeType string, data []byte, location CloudFileLocation) (file CloudFile, err error) {
	createFile, err := c.googleDriver.CreateFile(location.GoogleDriveUploadLocation(), name, data, mimeType)
	if err != nil {
		return
	}
	file = NewCloudFileFromGoogleDriveFile(createFile)
	return
}

func NewCloudFileFromGoogleDriveFile(dFile *drive.File) (file CloudFile) {
	file.Name = dFile.Name
	file.Location = CloudFileLocationObsidian
	return
}
