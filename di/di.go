package di

import (
	api "2023_asset_management/interface/rest_api"
	_ "github.com/google/subcommands"
	_ "github.com/google/wire"
)

type DI struct {
	RestAPI api.ServerInterface
}

func NewDI(restAPI api.ServerInterface) *DI {
	return &DI{RestAPI: restAPI}
}
