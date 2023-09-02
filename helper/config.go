package helper

import "github.com/spf13/viper"

var Config Configer

type Configer interface {
	Port() string
}

type config struct {
}

func (m *config) Port() string {
	return viper.GetString("PORT")
}

func newConfig() Configer {
	return &config{}
}

func init() {
	Config = newConfig()
}
