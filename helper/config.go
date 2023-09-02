package helper

import "github.com/spf13/viper"

var PORT = func() string {
	return viper.GetString("PORT")
}
