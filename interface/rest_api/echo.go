package api

import (
	"2023_asset_management/application"
	"2023_asset_management/domain"
	"bytes"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"mime/multipart"
	"net/http"
)

type EchoServer struct {
	fileStorer application.FileStorer
}

func (e *EchoServer) V1RedirectToPublicLink(ctx echo.Context, location V1AssetLocation, fileName string) error {
	link, err := e.fileStorer.GetPreviewLinkByLocationAndFileName(location.DomainLocation(), fileName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return ctx.Redirect(http.StatusFound, link)
}

func NewEchoServer(fileStorer application.FileStorer) ServerInterface {
	return &EchoServer{fileStorer: fileStorer}
}

func (location V1AssetLocation) IsValid() bool {
	switch location {
	case V1AssetLocationBlog, V1AssetLocationObsidian:
		return true
	}
	return false
}
func (location V1AssetLocation) DomainLocation() domain.CloudFileLocation {
	switch location {
	case V1AssetLocationObsidian:
		return domain.CloudFileLocationObsidian
	case V1AssetLocationBlog:
		return domain.CloudFileLocationBlog
	}
	return domain.CloudFileLocationObsidian
}

func (e *EchoServer) V1UploadAsset(ctx echo.Context, location V1AssetLocation) error {
	if !location.IsValid() {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("location is invalid"))
	}
	file, err := ctx.FormFile("image")
	open, err := file.Open()

	defer func(open multipart.File) {
		err := open.Close()
		if err != nil {
			fmt.Println("don't close the file open")
		}
	}(open)
	fileBytes := bytes.NewBuffer(nil)
	_, err = io.Copy(fileBytes, open)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	asset, err := e.fileStorer.UploadAsset(file.Filename, fileBytes.Bytes(), location.DomainLocation(), "image/png")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	link, err := e.fileStorer.GetPreviewLink(asset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, TempAsset{
		Url: link,
	})
}
