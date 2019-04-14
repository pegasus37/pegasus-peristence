package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

func main() {
	router := mux.NewRouter()

	setUpEnv()
	setUpConfig()
	setUpRouter(router)

	startServer(router)
}

func setUpEnv() {
	if os.Getenv("env") == "" {
		os.Setenv("env", "local")
	}
}

func setUpConfig() {
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./../config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
}

func startServer(router *mux.Router) {
	fmt.Println(viper.GetString("app.name") + " is running on port: " + viper.GetString("app.port"))
	fmt.Println("env: " + os.Getenv("env"))

	err := http.ListenAndServe(":"+viper.GetString("app.port"), router)
	if err != nil {
		log.Fatal("ListenAndServe Error:", err)
	}
}
