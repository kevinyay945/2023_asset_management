package api

import (
	"2023_asset_management/application"
	"2023_asset_management/domain"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"net/http"
	"os"

	"github.com/gavv/httpexpect/v2"
	"net/http/httptest"
	"testing"
)

type Suite struct {
	suite.Suite
	mockCtrl   *gomock.Controller
	request    *httpexpect.Expect
	server     *httptest.Server
	fileStorer *application.MockFileStorer
}

func TestSuiteInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (t *Suite) SetupTest() {
	t.mockCtrl = gomock.NewController(t.Suite.T())
	e := echo.New()
	t.fileStorer = application.NewMockFileStorer(t.mockCtrl)
	server := NewEchoServer(t.fileStorer)
	RegisterHandlers(e.Group(""), server)

	t.server = httptest.NewServer(e)

	t.request = httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  t.server.URL,
		Reporter: httpexpect.NewAssertReporter(t.T()),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t.T(), false),
		},
	})
}

func (t *Suite) TearDownTest() {
	defer t.mockCtrl.Finish()
	defer t.server.Close()
}

func (t *Suite) Test_V1_upload_asset_success() {
	data, _ := os.ReadFile("./wakuwaku.jpeg")
	file := domain.CloudFile{}
	t.fileStorer.EXPECT().
		UploadAsset("wakuwaku.jpeg", data, "obsidian").
		Return(file, nil)
	t.fileStorer.EXPECT().GetPreviewLink(file).Return("http://localhost/link", nil)

	resp := t.request.POST("/v1/asset/obsidian").
		WithMultipart().
		WithFileBytes("image", "wakuwaku.jpeg", data).
		Expect()

	resp.Status(http.StatusOK)
	resp.JSON().Object().Value("url").IsEqual("http://localhost/link")
}

func (t *Suite) Test_V1_upload_asset_path_invalid() {
	t.fileStorer.EXPECT().
		UploadAsset(gomock.Any(), gomock.Any(), gomock.Any()).
		Times(0).
		Return(domain.CloudFile{}, nil)
	t.fileStorer.EXPECT().GetPreviewLink(domain.CloudFile{}).Times(0).Return("http://localhost/link", nil)
	resp := t.request.POST("/v1/asset/invalid_path").Expect()

	resp.Status(http.StatusBadRequest)
}
