package config

import (
	"log"

	"github.com/spf13/viper"
)

func getEnviroment(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}

type Enviroment struct {
	Port     string
	Stage    string
	Database string
	Username string
	Password string
	Sslmode  string
}

func EnvData() Enviroment {
	env := Enviroment{
		Port:     getEnviroment("PORT"),
		Stage:    getEnviroment("STAGE"),
		Database: getEnviroment("DATABASE"),
		Username: getEnviroment("USERNAME"),
		Password: getEnviroment("PASSWORD"),
		Sslmode:  getEnviroment("SSLMODE"),
	}

	return env
}
