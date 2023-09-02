package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"net/http"

	"net/http/httptest"
	"testing"
)

type Suite struct {
	suite.Suite
	mockCtrl *gomock.Controller
}

func TestSuiteInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (t *Suite) SetupTest() {
	t.mockCtrl = gomock.NewController(t.Suite.T())
}

func (t *Suite) TearDownTest() {
	defer t.mockCtrl.Finish()
}

func (t *Suite) Test_V1_upload_asset_success() {
	e := echo.New()
	server := NewEchoServer()
	req := httptest.NewRequest(http.MethodPost,
		"/",
		nil,
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := server.V1UploadAsset(c)
	t.NoError(err)
	t.Equal(http.StatusOK, rec.Code)

	expectString, _ := t.jsonBytesParser([]byte((`{"url":"http://localhost/link"}`)))
	actualString, actualMap := t.jsonBytesParser(rec.Body.Bytes())
	t.Equal(expectString, actualString)
	t.Equal("http://localhost/link", actualMap["url"])
}

func (t *Suite) jsonBytesParser(ex []byte) (string, map[string]interface{}) {
	var expectOutput map[string]interface{}
	err := json.Unmarshal(ex, &expectOutput)
	t.NoError(err)
	expectOutputJson, _ := json.MarshalIndent(expectOutput, "", "  ")
	expected := string(expectOutputJson)
	return expected, expectOutput
}
