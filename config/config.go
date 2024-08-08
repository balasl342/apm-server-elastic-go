package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Database struct {
		DBHost     string `json:"dbhost"`
		DBPort     string `json:"dbport"`
		DBUser     string `json:"dbuser"`
		DBPassword string `json:"dbpassword"`
		DBName     string `json:"dbname"`
	} `json:"database"`
	ElasticAPM struct {
		APMServerURLs  string `json:"apmserverurls"`
		APMServiceName string `json:"apmservicename"`
		APMSecretToken string `json:"apmsecrettoken"`
		APMEnvironment string `json:"apmenvironment"`
	} `json:"elastic_apm"`
}

var AppConfig Config

func LoadConfig() error {
	file, err := os.Open("config/config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		return err
	}

	return nil
}
