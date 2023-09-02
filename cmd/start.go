/*
Copyright © 2023 Kevin Chen
*/
package cmd

import (
	"2023_asset_management/asset"
	"2023_asset_management/di"
	"2023_asset_management/helper"
	"crypto/subtle"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start main program",
	Long:  `start main program`,
	Run: func(cmd *cobra.Command, args []string) {
		diObj := di.InitializeDICmd()
		go StartRestAPI(diObj)
		select {}
	},
}

func StartRestAPI(*di.DI) {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Skipper: func(c echo.Context) bool {
			if c.Request().RequestURI == "/" {
				return true
			}
			if strings.HasPrefix(c.Request().RequestURI, "/doc") {
				return false
			}
			return true
		},
		Validator: func(username, password string, c echo.Context) (bool, error) {
			docusr := helper.Config.DocUser()
			docpwd := helper.Config.DocPwd()
			if subtle.ConstantTimeCompare([]byte(username), []byte(docusr)) == 1 &&
				subtle.ConstantTimeCompare([]byte(password), []byte(docpwd)) == 1 {
				return true, nil
			}
			return false, nil
		},
	}))
	e.FileFS("/doc/api", "index.html", echo.MustSubFS(asset.IndexHTML, "swagger"))
	e.StaticFS("/doc/api", echo.MustSubFS(asset.Dist, "swagger"))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", helper.Config.Port())))
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
