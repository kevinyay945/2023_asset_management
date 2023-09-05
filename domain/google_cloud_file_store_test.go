package domain

import (
	"2023_asset_management/infrastructure/googledrive"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"google.golang.org/api/drive/v3"
	"os"
	"testing"
)

type GoogleCloudFileStoreSuite struct {
	suite.Suite
	mockCtrl     *gomock.Controller
	store        CloudFileStorer
	googleDriver *googledrive.MockGoogleDriver
}

func TestSuiteInitGoogleCloudFileStore(t *testing.T) {
	suite.Run(t, new(GoogleCloudFileStoreSuite))
}

func (t *GoogleCloudFileStoreSuite) SetupTest() {
	t.mockCtrl = gomock.NewController(t.Suite.T())
	t.googleDriver = googledrive.NewMockGoogleDriver(t.mockCtrl)
	t.store = NewGoogleCloudFileStore(t.googleDriver)
}

func (t *GoogleCloudFileStoreSuite) TearDownTest() {
	defer t.mockCtrl.Finish()
}

func (t *GoogleCloudFileStoreSuite) Test_upload_file() {
	data, err := os.ReadFile("../asset/test/wakuwaku.jpeg")
	t.NoError(err)

	file := drive.File{Name: "wakuwaku.jpeg"}
	t.googleDriver.EXPECT().CreateFile(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&file, nil)

	_, err = t.store.UploadFile("wakuwaku.jpeg", "image/png", data, CloudFileLocationObsidian)
	t.NoError(err)
}

func (t *GoogleCloudFileStoreSuite) Test_upload_file_check_for_invalid_mime_type() {
	data, err := os.ReadFile("../asset/test/wakuwaku.jpeg")
	t.NoError(err)
	file := drive.File{Name: "wakuwaku.jpeg"}

	t.googleDriver.EXPECT().
		CreateFile(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Times(0).
		Return(&file, nil)

	_, err = t.store.UploadFile("wakuwaku.jpeg", "invalid mime type", data, CloudFileLocationObsidian)
	t.Error(err)
}
