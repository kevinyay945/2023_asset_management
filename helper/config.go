package helper

import "github.com/spf13/viper"

var Config Configer

type Configer interface {
	Port() int
	DocUser() string
	DocPwd() string
}

type config struct {
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
