package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Manager struct{}

func NewManager(configPath string,credsPath string,  configName string) Manager {
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(credsPath)
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	log.Info(fmt.Sprintf("Got the following config: %v", viper.AllSettings()))
	return Manager{}
}

func (Manager) GetInt(value string) int {
	return viper.GetInt(value)
}

func (Manager) GetString(value string) string {
	return viper.GetString(value)
}

func (Manager) GetBool(value string) bool {
	return viper.GetBool(value)
}
