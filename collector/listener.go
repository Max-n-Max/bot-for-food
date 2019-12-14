package collector

import (
	"context"
	"fmt"
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"github.com/bitfinexcom/bitfinex-api-go/v2/websocket"
	"log"
	"time"
)

func startOrderBookListener(wClient *websocket.Client, fn queryType, pair, precision, frequency string, priceLevel int, channel chan string) (string, error) {
	err := wClient.Connect()
	if err != nil {
		log.Printf("could not connect: %s", err.Error())
		return "", err
	}

	id, err := wClient.
		SubscribeBook(context.Background(),
			bitfinex.TradingPrefix+ pair,
			bitfinex.BookPrecision(precision), // P0-3 | R0
			bitfinex.BookFrequency(frequency), // F1  print in RT or by junks
			priceLevel)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	go listen(wClient, fn, pair, channel)

	return id, nil
}

func listen(wClient *websocket.Client, fn queryType, pair string, channel chan string) {
	lastTime := time.Now()
	for obj := range wClient.Listen() {

		switch obj.(type) {
		case *bitfinex.BookUpdateSnapshot:
			//log.Printf("BookUpdateSnapshot: %#v", obj)
			break
		case *bitfinex.BookUpdate:
			now := time.Now()
			if now.Sub(lastTime) < 10 * time.Second {
				break
			}
			lastTime = now

			bookUpdate := obj.(*bitfinex.BookUpdate)
			side := "unknown"
			if bookUpdate.Side == 0x1 {
				side = "BUY "
			} else if bookUpdate.Side == 0x2 {
				side = "SELL"
			}
			log.Printf("BookUpdate: side=%s price= %#f; amount= %#f;", side, bookUpdate.Price, bookUpdate.Amount)

			res, e := fn(pair)
			if e == nil {
				channel <- res
			} else {
				fmt.Println("Cannot get data from exchange. Error: ", e)
			}
			break
		}
	}
}
