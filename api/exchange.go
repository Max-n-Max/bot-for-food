package api

import (
	"github.com/bitfinexcom/bitfinex-api-go/v1"
)

type BitfinexWrapper struct {
	api              *bitfinex.Client
}

type ExchangeWrapper interface {

}

func NewBitfinexWrapper(publicKey string, secretKey string) ExchangeWrapper {
	return &BitfinexWrapper{
		api:                 bitfinex.NewClient().Auth(publicKey, secretKey),
	}
}