package main

import (
	"fmt"
	v1 "github.com/bitfinexcom/bitfinex-api-go/v1"
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	v2 "github.com/bitfinexcom/bitfinex-api-go/v2/rest"
	"testing"
)

func TestBitfinexGetOrderBook(t *testing.T) {
	b := v1.NewClient()
	b = b.Auth("22dsdssddssd","2dssdsdsdsdsd22")

	orders, err := b.OrderBook.Get(bitfinex.BTCUSD, 10, 10, false)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(orders)
}

func TestBitfinexGetTrades(t *testing.T) {
	b := v2.NewClient()
	//b = b.Auth("","")

	candles, err := b.Candles.History(bitfinex.TradingPrefix + bitfinex.BTCUSD, bitfinex.OneMinute)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(candles)
}
