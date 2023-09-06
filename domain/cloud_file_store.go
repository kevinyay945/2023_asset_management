package domain

import "2023_asset_management/infrastructure/googledrive"

type CloudFileLocation int

func (l CloudFileLocation) GoogleDriveUploadLocation() googledrive.GoogleDriveUploadLocation {
	switch l {
	case CloudFileLocationObsidian:
		return googledrive.GoogleDriveUploadLocationObsidian
	case CloudFileLocationBlog:
		return googledrive.GoogleDriveUploadLocationBlog
	}
	return 0
}

func (l CloudFileLocation) LinkAlias() string {
	switch l {
	case CloudFileLocationObsidian:
		return "obsidian"
	case CloudFileLocationBlog:
		return "blog"
	}
	return ""
}

const (
	_ CloudFileLocation = iota
	CloudFileLocationObsidian
	CloudFileLocationBlog
)

type CloudFile struct {
	Name     string
	Location CloudFileLocation
}

var validMimeType = []string{"image/png"}

//go:generate mockgen -destination=cloud_file_store.mock.go -package=domain -self_package=2023_asset_management/domain . CloudFileStorer
type CloudFileStorer interface {
	UploadFile(name string, mimeType string, data []byte, location CloudFileLocation) (file CloudFile, err error)
	GetPublicLink(file CloudFile) (link string, err error)
	GetCloudFileByNameAndLocation(name string, location CloudFileLocation) (file CloudFile, err error)
}
