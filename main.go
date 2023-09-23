package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Println("couldn’t find .env file. reading from env vars.")
	}
}

func main() {
	port := viper.GetString("PORT")
	host := viper.GetString("HOST")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	e.HideBanner = true
	e.Logger.Fatal(e.Start(fmt.Sprintf("%v:%v", host, port)))
}
