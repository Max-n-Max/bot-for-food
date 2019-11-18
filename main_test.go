package main

import (
	"context"
	"fmt"
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	v2 "github.com/bitfinexcom/bitfinex-api-go/v2/rest"

	"github.com/bitfinexcom/bitfinex-api-go/v2/websocket"
	"log"
	"testing"
	"time"
)

/**
https://medium.com/bitfinex/tutorial-trading-cryptos-on-the-bitfinex-platform-using-golang-9b100ddcf72c
*/
func TestListenOnChage(t *testing.T){

	client := websocket.New()
	err := client.Connect()
	if err != nil {
		log.Printf("could not connect: %s", err.Error())
		return
	}

	_, err = client.SubscribeTrades(context.Background(), bitfinex.TradingPrefix+bitfinex.ETHBTC)
	if err != nil {
		log.Fatal(err)
	}

	for obj := range client.Listen() {
		switch obj.(type) {
		case error:
			log.Printf("EROR RECV: %s", obj)
		default:
			log.Printf("MSG RECV: %#v", obj)
		}
	}
}

func TestCandlesHistoryWithQuery(t *testing.T){

}


func TestCandlesHistory(t *testing.T){

	c := v2.NewClient()


	log.Printf("1) Query Last Candle")
	candle, err := c.Candles.Last(bitfinex.TradingPrefix+bitfinex.BTCUSD, bitfinex.FiveMinutes)

	if err != nil {
		log.Fatalf("getting candle: %s", err)
	}

	log.Printf("last candle: %#v\n", candle)

	now := time.Now()
	millis := now.UnixNano() / 1000000

	prior := now.Add(time.Duration(-24) * time.Hour)
	millisStart := prior.UnixNano() / 1000000


	log.Printf("2) Query Candle History with no params")
	candles, err := c.Candles.History(bitfinex.TradingPrefix+bitfinex.BTCUSD, bitfinex.FiveMinutes)

	if err != nil {
		log.Fatalf("getting candles: %s", err)
	}

	log.Printf("length of candles is: %v", len(candles.Snapshot))

	log.Printf("first candle is: %#v\n", candles.Snapshot[0])
	log.Printf("last candle is: %#v\n", candles.Snapshot[len(candles.Snapshot)-1])

	start := bitfinex.Mts(millisStart)
	end := bitfinex.Mts(millis)

	log.Printf("3) Query Candle History with params")
	candlesMore, err := c.Candles.HistoryWithQuery(
		bitfinex.TradingPrefix+bitfinex.BTCUSD,
		bitfinex.FiveMinutes,
		start,
		end,
		200,
		bitfinex.OldestFirst,
	)

	if err != nil {
		log.Fatalf("getting candles: %s", err)
	}

	log.Printf("length of candles is: %v", len(candlesMore.Snapshot))
	log.Printf("first candle is: %#v\n", candlesMore.Snapshot[0])
	log.Printf("last candle is: %#v\n", candlesMore.Snapshot[len(candlesMore.Snapshot)-1])


	//	//client := rest.NewClient()
//	client := bitfinex.NewClient()
//	candles, err := client.Candles.History(bfx.TradingPrefix+bfx.BTCUSD, bfx.FiveMinutes)
//	if err != nil {
//		log.Fatalf("Failed getting candles: %s", err)
//	}
//	log.Printf("length of candles is: %v", len(candles.Snapshot))
//	log.Printf("first candle is: %#v\n", candles.Snapshot[0])
//	log.Printf("last candle is: %#v\n", candles.Snapshot[len(candles.Snapshot)-1])
//}
//
//func TestBitfinexGetOrderBook(t *testing.T) {
//	b := v1.NewClient()
//	b = b.Auth("22dsdssddssd","2dssdsdsdsdsd22")
//
//	orders, err := b.OrderBook.Get(bitfinex.BTCUSD, 10, 10, false)
//	if err != nil {
//		fmt.Println("Error", err)
//	}
//	fmt.Println(orders)
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
	//m, _ := db.NewManager()
	////"{"timestamp":{$gte: "2019-11-13", $lt: "2019-11-14"}}"
	//
	//var results []resources.OrderBook
	//
	////r := record{Timestamp:timestamp{gte:"2019-11-13", lt:"2019-11-14"}}
	//col := m.GetDB().DB("cryptodb").C("orderbook")
	//_ = col.Find(bson.M{"timestamp": bson.M{"$gt": "2019-11-13", "$lt": "2019-11-14"}}).All(&results)
	//fmt.Println(results)
}
