package main

import (
	"fmt"
	"github.com/Max-n-Max/bot-for-food/cmd"
	"github.com/Max-n-Max/bot-for-food/db"
	v1 "github.com/bitfinexcom/bitfinex-api-go/v1"
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	v2 "github.com/bitfinexcom/bitfinex-api-go/v2/rest"
	"gopkg.in/mgo.v2/bson"
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


func TestMongo(t *testing.T) {
	m, _ := db.NewManager()
	//"{"timestamp":{$gte: "2019-11-13", $lt: "2019-11-14"}}"

	var results []cmd.OrderBook

	//r := record{Timestamp:timestamp{gte:"2019-11-13", lt:"2019-11-14"}}
	col := m.GetDB().DB("cryptodb").C("orderbook")
	_ = col.Find(bson.M{"timestamp": bson.M{"$gt": "2019-11-13", "$lt": "2019-11-14"}}).All(&results)
	fmt.Println(results)
}
