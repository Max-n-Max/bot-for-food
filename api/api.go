package api

import (
	"bot-for-food/env"
	"bot-for-food/exchanges"
)



func InitExchange(exchangeConfig env.ExchangeConfig) exchanges.ExchangeWrapper {

	var exch exchanges.ExchangeWrapper = exchanges.NewBitfinexWrapper(exchangeConfig)

	return exch
}