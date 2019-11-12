package api

import (
	"bot-for-food/config"
	"bot-for-food/exchanges"
)



func InitExchange(exchangeConfig config.ExchangeConfig) exchanges.ExchangeWrapper {

	var exch exchanges.ExchangeWrapper = exchanges.NewBitfinexWrapper(exchangeConfig)

	return exch
}