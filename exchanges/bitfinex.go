package exchanges

import (
	"bot-for-food/env"
	"github.com/bitfinexcom/bitfinex-api-go/v1"
)

type BitfinexWrapper struct {
	api              *bitfinex.Client
}

type ExchangeWrapper interface {

}

func NewBitfinexWrapper(exchangeConfig env.ExchangeConfig) ExchangeWrapper {
	return &BitfinexWrapper{
		api:                 bitfinex.NewClient().Auth(exchangeConfig.PublicKey, exchangeConfig.SecretKey),
	}
}