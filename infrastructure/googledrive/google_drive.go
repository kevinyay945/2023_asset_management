package googledrive

import (
	"2023_asset_management/helper"
	"bytes"
	"context"
	"fmt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"io"
	"slices"
)

type GoogleDriveUploadLocation int

type property struct {
	parentId string
}

// https://console.cloud.google.com/projectselector2/iam-admin/serviceaccounts
var googleDriveUploadLocationProperty map[GoogleDriveUploadLocation]property = map[GoogleDriveUploadLocation]property{
	GoogleDriveUploadLocationObsidian: {
		parentId: "12o01NrAVom6FSrQgKfQVsR2G5H0dVsd3",
	},
	GoogleDriveUploadLocationBlog: {
		parentId: "1zH0JPFAiBkqKWQPKTQ49L6xC2IuGCrny",
	},
}

func (l GoogleDriveUploadLocation) ParentID() (string, error) {
	p, ok := googleDriveUploadLocationProperty[l]
	if !ok {
		return "", newGoogleDriveError("GoogleDriveUploadLocation.ParentID", fmt.Errorf("Fail to get folderID"))
	}
	id := p.parentId
	if id == "" {
		return "", newGoogleDriveError("GoogleDriveUploadLocation.ParentID", fmt.Errorf("Fail to get folderID"))
	}
	return id, nil
}

const (
	_ GoogleDriveUploadLocation = iota
	GoogleDriveUploadLocationObsidian
	GoogleDriveUploadLocationBlog
)

type GoogleDriver interface {
}
type GoogleDrive struct {
	driveService *drive.Service
}

func NewGoogleDrive() *GoogleDrive {
	d := GoogleDrive{}
	token := helper.Config.GoogleDriveApiToken()
	d.setCredentialsJson(token)
	return &d
}

func (g *GoogleDrive) CheckAuthorization() (err error) {
	_, err = g.driveService.Files.List().PageSize(10).Fields("nextPageToken, files(id, name)").Do()
	return
}

func (g *GoogleDrive) setCredentialsJson(token string) {
	tokenByte := []byte(token)
	service, _ := drive.NewService(context.Background(), option.WithCredentialsJSON(tokenByte))
	g.driveService = service
}

func (g *GoogleDrive) GetFileParentId(location GoogleDriveUploadLocation, fileName string) ([]string, error) {

	query := fmt.Sprintf("name = '%s'", fileName)
	r, err := g.driveService.Files.List().PageSize(10).Q(query).
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		return nil, newGoogleDriveError("GetFileParentId", err)
	}
	for _, i := range r.Files {
		fmt.Printf("%s (%s)\n", i.Name, i.Id)
		get := g.driveService.Files.Get(i.Id)
		get.Fields("parents")
		do, err := get.Do()
		if err != nil {
			return nil, newGoogleDriveError("GetFileParentId", err)
		}
		if len(do.Parents) >= 1 {
			return do.Parents, nil
		}
	}
	return nil, newGoogleDriveError("GetFileParentId", fmt.Errorf("file not found in google drive"))
}

func (g *GoogleDrive) CreateFile(location GoogleDriveUploadLocation, fileName string, data []byte, mimeType string) (file *drive.File, err error) {

	mimeTypeAllowType := []string{"image/png"}
	if !slices.Contains(mimeTypeAllowType, mimeType) {
		err = newGoogleDriveError("CreateFile", fmt.Errorf("mimeType is only support for %v", mimeTypeAllowType))
		return
	}
	folderId, err := location.ParentID()
	if err != nil {
		err = newGoogleDriveError("CreateFile", err)
		return
	}
	buffer := bytes.NewBuffer(data)
	file, err = createFile(
		g.driveService, fileName, mimeType, buffer, folderId,
	)
	if err != nil {
		err = newGoogleDriveError("CreateFile", err)
		return
	}
	return
}

func (g *GoogleDrive) GetFilePublicLink(blog GoogleDriveUploadLocation, fileName string) (link string, err error) {
	parentId, err := blog.ParentID()
	if err != nil {
		return "", newGoogleDriveError("GetFilePublicLink", err)
	}
	query := fmt.Sprintf("name = '%s' and '%s' in parents", fileName, parentId)
	r, err := g.driveService.Files.List().PageSize(10).Q(query).
		Fields("nextPageToken, files(id, name,parents)").Do()
	if err != nil {
		err = newGoogleDriveError("GetFilePublicLink", err)
		return
	}
	if len(r.Files) == 0 {
		err = newGoogleDriveError("GetFilePublicLink", fmt.Errorf("File not found in google drive"))
		return
	} else {
		for _, i := range r.Files {
			if err != nil {
				err = newGoogleDriveError("GetFilePublicLink", err)
				return
			}
			link = fmt.Sprintf("https://drive.google.com/uc?export=view&id=%s\n", i.Id)
		}
	}
	return
}
func createFile(service *drive.Service, name string, mimeType string, content io.Reader, parentID string) (*drive.File, error) {
	parents := []string{}
	if parentID != "" {
		parents = []string{parentID}
	}
	f := &drive.File{
		MimeType: mimeType,
		Name:     name,
		Parents:  parents,
	}
	file, err := service.Files.Create(f).Media(content).Do()

	if err != nil {
		return nil, err
	}

	return file, nil
}
