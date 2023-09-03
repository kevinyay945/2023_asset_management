package helper

import "github.com/spf13/viper"

var Config Configer

//go:generate mockgen -destination=config.mock.go -package=helper -self_package=2023_asset_management/helper . Configer
type Configer interface {
	Port() int
	DocUser() string
	DocPwd() string
	GoogleDriveApiToken() string
}

type config struct {
}

func (m *config) GoogleDriveApiToken() string {
	return viper.GetString("GOOGLE_DRIVE_API_TOKEN")
}

func (m *config) DocUser() string {
	return viper.GetString("DOCUSR")
}

func (m *config) DocPwd() string {
	return viper.GetString("DOCPWD")
}

func (m *config) Port() int {
	return viper.GetInt("PORT")
}

func newConfig() Configer {
	return &config{}
}

func init() {
	Config = newConfig()
}
