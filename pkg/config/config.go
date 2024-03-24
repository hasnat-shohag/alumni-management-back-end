package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DbUser           string `mapstructure:"DBUSER"`
	DbPass           string `mapstructure:"DBPASS"`
	DbIp             string `mapstructure:"DBIP"`
	DbName           string `mapstructure:"DBNAME"`
	Port             string `mapstructure:"PORT"`
	JwtSecret        string `mapstructure:"JWT_SECRET"`
	JwtExpireMinutes int    `mapstructure:"JWT_EXPIRE_MINUTES"`
}

var LocalConfig *Config

func initConfig() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file ", err)
	}

	var config *Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error reading env file", err)
	}

	return config
}

func SetConfig() {
	LocalConfig = initConfig()
}
