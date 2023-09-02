package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type EchoServer struct{}

func NewEchoServer() ServerInterface {
	return &EchoServer{}
}

func (e *EchoServer) V1UploadAsset(ctx echo.Context) error {
	//TODO implement me
	return ctx.JSON(http.StatusOK, "v1UploadAsset")
}
