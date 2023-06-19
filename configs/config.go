package configs

import (
	"log"

	"github.com/spf13/viper"
)

var (
	Conf *Config
)

type Config struct {
	App struct {
		Name   string
		Server struct {
			Port string `json:"port"`
		} `json:"server"`
	} `json:"app"`
}

func LoadConfig() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("unable to read config, %v", err)
	}

	err = viper.Unmarshal(&Conf)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
