package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ExchangeConfig struct {
	ExchangeName string `json:"name"`
	PublicKey    string `json:"public_key"`
	SecretKey    string `json:"secret_key"`
}

type BotConfig struct {
	ExchangeConfigs  []ExchangeConfig `json:"exchange_configs"`
}


//Configuration Global config declaration for wide App use.
var Configuration = &BotConfig{}


func Init()  (*BotConfig, error) {
	err := marshalConfig()
	if err != nil{
		return nil, err
	}
	return Configuration, nil
}

//getConfigPath get configuration file path
func getConfigPath() string {
	confPath := "config/"

	switch {
	case isArgTrue(os.Getenv("TEST_ENV")):
		confPath += "test-config.json"
	default:
		panic("Error: Use environment variables (PROD_ENV/DEV_ENV/TEST_ENV) to load config file.")
	}
	return confPath
}

func isArgTrue(arg string) bool{
  return arg == "true"
}

func marshalConfig() error {
	file, err := ioutil.ReadFile(getConfigPath())
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, Configuration)
	if err != nil {
		return err
	}
	fmt.Printf("appconf: %+v", Configuration)
	return nil
}