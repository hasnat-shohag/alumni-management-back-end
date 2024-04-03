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

// SMTPConfig is a struct that defines the configuration for the SMTP server.
type SMTPConfig struct {
	Host       string `mapstructure:"SMTP_HOST"`
	Port       string `mapstructure:"SMTP_PORT"`
	Username   string `mapstructure:"SMTP_USERNAME"`
	Password   string `mapstructure:"SMTP_PASSWORD"`
	AdminEmail string `mapstructure:"ADMIN_EMAIL"`
}

// InitSMTPConfig initializes the SMTP configuration.
func InitSMTPConfig() *SMTPConfig {
	viper.AddConfigPath(".")
	viper.SetConfigName("base")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file ", err)
	}

	var smtpConfig *SMTPConfig
	if err := viper.Unmarshal(&smtpConfig); err != nil {
		log.Fatal("Error reading env file", err)
	}

	return smtpConfig
}

var LocalSMTPConfig = InitSMTPConfig()
