package cmd

import (
	"bot-for-food/api"
	"bot-for-food/config"
	"bot-for-food/exchanges"
	"fmt"
)


var botConfig config.BotConfig

func Execute() {
	fmt.Print("Getting configurations ... ")

	botConfig, err := config.Init()
	if err != nil {
		fmt.Println("Cannot read from configuration file")
		return
	}
	fmt.Println("DONE")

	fmt.Print("Getting exchange info ... ")
	wrappers := make([]exchanges.ExchangeWrapper, len(botConfig.ExchangeConfigs))
	for i, config := range botConfig.ExchangeConfigs {
		wrappers[i] = api.InitExchange(config)
	}
	fmt.Println("DONE")
}