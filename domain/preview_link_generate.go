package domain

import (
	"2023_asset_management/helper"
	"fmt"
)

//go:generate mockgen -destination=preview_link_generate.mock.go -package=domain -self_package=2023_asset_management/domain . PreviewLinkGenerator
type PreviewLinkGenerator interface {
	GetLinkByCloudFile(file CloudFile) string
}

type PreviewLinkGenerate struct {
}

func NewPreviewLinkGenerate() PreviewLinkGenerator {
	return &PreviewLinkGenerate{}
}

func (p *PreviewLinkGenerate) GetLinkByCloudFile(file CloudFile) string {
	return fmt.Sprintf("%s/v1/temp-link/%s/%s", helper.Config.BaseURL(), file.Location.LinkAlias(), file.Name)
}
