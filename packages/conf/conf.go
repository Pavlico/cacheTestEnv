package conf

import (
	"log"

	"github.com/spf13/viper"
)

const NotFoundMessage = "User doesn't exist"
const UserIdKey = "id"
const configFileType = "json"
const configPath = "."
const configFileName = "db"
const configDbUsername = "Username"
const configDbPassword = "Password"
const configDbHostname = "Hostname"
const configDbName = "DbName"

type dbCredentials struct {
	Username string
	Password string
	Hostname string
	DbName   string
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName(configFileName)
	v.SetConfigType(configFileType)
	err := v.ReadInConfig()
	if err != nil {
		return v, err
	}
	return v, nil
}

func GetDbCredentials() *dbCredentials {
	config, err := LoadConfig()
	if err != nil {
		log.Println(err)
		return nil
	}
	return &dbCredentials{
		Username: config.GetString(configDbUsername),
		Password: config.GetString(configDbPassword),
		Hostname: config.GetString(configDbHostname),
		DbName:   config.GetString(configDbName),
	}
}
