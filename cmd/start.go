package cmd

import (
	"bot-for-food/api"
	"bot-for-food/env"
	"bot-for-food/exchanges"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts trading using saved configs",
	Long:  `Starts trading using saved configs`,
	Run:   executeStartCommand,
}


var botConfig env.BotConfig

func init() {
	RootCmd.AddCommand(startCmd)

	//sayhelloCmd.Flags().StringP("name", "n", "", "Set your name")
	//hello-cli say hello -n Nick
}

func initConfigs() error {
	configFile, err := os.Open("./.bot_config.yml")
	if err != nil {
		return err
	}
	contentToMarshal, err := ioutil.ReadAll(configFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(contentToMarshal, &botConfig)
	if err != nil {
		return err
	}
	return nil
}

func executeStartCommand(cmd *cobra.Command, args []string) {
	fmt.Print("Getting configurations ... ")
	if err := initConfigs(); err != nil {
		fmt.Println("Cannot read from configuration file, please create or replace the current one using gobot init")
		return
	}
	fmt.Println("DONE")

	fmt.Print("Getting exchange info ... ")
	wrappers := make([]exchanges.ExchangeWrapper, len(botConfig.ExchangeConfigs))
	for i, config := range botConfig.ExchangeConfigs {
		wrappers[i] = api.InitExchange(config)
	}
	fmt.Println("DONE")

	fmt.Print("Getting markets cold info ... ")
	//for _, strategyConf := range botConfig.Strategies {
	//	mkts := make([]*environment.Market, len(strategyConf.Markets))
	//	for i, mkt := range strategyConf.Markets {
	//		currencies := strings.SplitN(mkt.Name, "-", 2)
	//		mkts[i] = &environment.Market{
	//			Name:           mkt.Name,
	//			BaseCurrency:   currencies[0],
	//			MarketCurrency: currencies[1],
	//		}
	//
	//		mkts[i].ExchangeNames = make(map[string]string, len(wrappers))
	//
	//		for _, exName := range mkt.Exchanges {
	//			mkts[i].ExchangeNames[exName.Name] = exName.MarketName
	//		}
	//	}
	//	err := strategies.MatchWithMarkets(strategyConf.Strategy, mkts)
	//	if err != nil {
	//		fmt.Println("Cannot add tactic : ", err)
	//	}
	//}
	fmt.Println("DONE")
}