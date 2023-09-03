package domain

type CloudFileLocation int

const (
	_ CloudFileLocation = iota
	CloudFileLocationObsidian
	CloudFileLocationBlog
)

type CloudFile struct{}

//go:generate mockgen -destination=cloud_file_store.mock.go -package=domain -self_package=2023_asset_management/domain . CloudFileStorer
type CloudFileStorer interface {
	UploadFile(name string, mimeType string, data []byte, location CloudFileLocation) (file CloudFile, err error)
	GetPublicLink(file CloudFile) (link string, err error)
}
