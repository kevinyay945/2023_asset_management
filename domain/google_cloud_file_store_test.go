package domain

import (
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"os"
	"testing"
)

type GoogleCloudFileStoreSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
	store    CloudFileStorer
}

func TestSuiteInitGoogleCloudFileStore(t *testing.T) {
	suite.Run(t, new(GoogleCloudFileStoreSuite))
}

func (t *GoogleCloudFileStoreSuite) SetupTest() {
	t.mockCtrl = gomock.NewController(t.Suite.T())
	t.store = NewGoogleCloudFileStore()
}

func (t *GoogleCloudFileStoreSuite) TearDownTest() {
	defer t.mockCtrl.Finish()
}

func (t *GoogleCloudFileStoreSuite) Test_upload_file() {
	data, err := os.ReadFile("../asset/test/wakuwaku.jpeg")
	t.NoError(err)
	_, err = t.store.UploadFile("wakuwaku.jpeg", "image/png", data, CloudFileLocationObsidian)
	t.NoError(err)
}
