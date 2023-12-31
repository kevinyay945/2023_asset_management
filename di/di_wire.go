//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package di

import (
	"2023_asset_management/application"
	"2023_asset_management/domain"
	"2023_asset_management/infrastructure/googledrive"
	api "2023_asset_management/interface/rest_api"
	"github.com/google/wire"
)

// InitializeAuthCmd creates an Auth Init Struct. It will error if the Event is staffed with
// a grumpy greeter.
func InitializeDICmd() *DI {
	wire.Build(
		googledrive.NewGoogleDrive,
		application.NewFileStore,
		domain.NewGoogleCloudFileStore,
		domain.NewPreviewLinkGenerate,
		api.NewEchoServer,
		NewDI,
	)
	return nil
}
