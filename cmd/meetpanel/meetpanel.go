package main

import (
	"fmt"
	"log"
	router "meetpanel/internal/router"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("../../")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Println("couldn’t find .env file. reading from env vars.")
	}
}

func main() {
	port := viper.GetString("PORT")
	host := viper.GetString("HOST")

	router := router.NewRouter()
	router.Logger.Fatal(router.Start(fmt.Sprintf("%v:%v", host, port)))
}
