package exchanges

import (
	"bot-for-food/config"
	"github.com/bitfinexcom/bitfinex-api-go/v1"
)

type BitfinexWrapper struct {
	api              *bitfinex.Client
}

type ExchangeWrapper interface {

}

func NewBitfinexWrapper(exchangeConfig config.ExchangeConfig) ExchangeWrapper {
	return &BitfinexWrapper{
		api:                 bitfinex.NewClient().Auth(exchangeConfig.PublicKey, exchangeConfig.SecretKey),
	}
}