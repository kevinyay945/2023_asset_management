package googledrive

import (
	"2023_asset_management/helper"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

type GoogleDriveSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
	drive    *GoogleDrive
}

func TestSuiteInitGoogleDrive(t *testing.T) {
	_, err := os.ReadFile("../../.config.dev.yaml")
	if err != nil {
		t.Skip("Skipping testing in production")
	}
	suite.Run(t, new(GoogleDriveSuite))
}

func (t *GoogleDriveSuite) SetupTest() {
	t.mockCtrl = gomock.NewController(t.Suite.T())
	mockConfiger := helper.NewMockConfiger(t.mockCtrl)
	configData, _ := os.ReadFile("../../.config.dev.yaml")
	credential := struct {
		Token string `yaml:"GOOGLE_DRIVE_API_TOKEN"`
	}{}
	err := yaml.Unmarshal(configData, &credential)
	t.NoError(err, "Fail to parse config file")
	t.NotEmpty(credential.Token, "GCP Token is empty")

	mockConfiger.EXPECT().GoogleDriveApiToken().Return(string(credential.Token))
	helper.Config = mockConfiger
	t.drive = NewGoogleDrive()
}

func (t *GoogleDriveSuite) TearDownTest() {
	defer t.mockCtrl.Finish()
}

func (t *GoogleDriveSuite) Test_check_token() {
	err := t.drive.CheckAuthorization()
	t.NoError(err)
}

func (t *GoogleDriveSuite) Test_get_parent_id() {
	id, err := t.drive.GetFileParentId(GoogleDriveUploadLocationBlog, "wakuwaku.png")
	t.NoError(err)
	t.Equal("test", id)
}

func (t *GoogleDriveSuite) Test_upload_file() {
	data, _ := os.ReadFile("../../asset/test/wakuwaku.jpeg")
	_, err := t.drive.CreateFile(GoogleDriveUploadLocationBlog, "wakuwaku.png", data, "image/png")
	t.NoError(err)
}

func (t *GoogleDriveSuite) Test_get_share_link() {
	link, err := t.drive.GetFilePublicLink(GoogleDriveUploadLocationBlog, "wakuwaku.png")
	t.NoError(err)
	t.Equal("test", link)
}
