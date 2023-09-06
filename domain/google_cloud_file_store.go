package domain

import (
	"2023_asset_management/infrastructure/googledrive"
	"fmt"
	"google.golang.org/api/drive/v3"
	"slices"
)

type GoogleCloudFileStore struct {
	googleDriver googledrive.GoogleDriver
}

func NewGoogleCloudFileStore(googleDriver googledrive.GoogleDriver) CloudFileStorer {
	return &GoogleCloudFileStore{googleDriver: googleDriver}
}

func (c *GoogleCloudFileStore) GetCloudFileByNameAndLocation(name string, location CloudFileLocation) (file CloudFile, err error) {
	file.Name = name
	file.Location = location
	return
}

func (c *GoogleCloudFileStore) GetPublicLink(file CloudFile) (link string, err error) {
	link, err = c.googleDriver.GetFilePublicLink(file.Location.GoogleDriveUploadLocation(), file.Name)
	return
}

func (c *GoogleCloudFileStore) UploadFile(name string, mimeType string, data []byte, location CloudFileLocation) (file CloudFile, err error) {
	if !slices.Contains(validMimeType, mimeType) {
		err = newGoogleCloudFileStoreError("UploadFile", fmt.Errorf("invalid mime type: %s", mimeType))
		return
	}
	createFile, err := c.googleDriver.CreateFile(location.GoogleDriveUploadLocation(), name, data, mimeType)
	if err != nil {
		return
	}
	file = NewCloudFileFromGoogleDriveFile(createFile, location)
	return
}

func NewCloudFileFromGoogleDriveFile(dFile *drive.File, location CloudFileLocation) (file CloudFile) {
	file.Name = dFile.Name
	file.Location = location
	return
}
