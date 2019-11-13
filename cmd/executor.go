package cmd

import (
	"fmt"
	"github.com/Max-n-Max/bot-for-food/collector"
	"github.com/Max-n-Max/bot-for-food/db"
)

type Order struct {
	Price     string `json:"Price"`
	Rate      string `json:"Rate"`
	Amount    string `json:"Amount"`
	Period    int    `json:"Period"`
	Timestamp string `json:"Timestamp"`
	Frr       string `json:"Frr"`
}

type OrderBook struct {
	Timestamp string
	BidsWall  float64
	AsksWall  float64
	Window    float64
	Bids      []Order
	Asks      []Order

}

func Execute(collector collector.Manager, db db.Manager) {
	orderBookExchangeCh := make(chan string)
	tradesExchangeCh := make(chan string)

	orderBookProcessCh := make(chan OrderBook)
	//tradesProcessCh := make(chan OrderBook)
	
	go collector.RunOrderBookTicker(orderBookExchangeCh)
	go collector.RunTradesTicker(tradesExchangeCh)

	go ProcessOrderBook(orderBookExchangeCh, orderBookProcessCh)
	go saveResults(db, orderBookProcessCh, "orderbook")
}

func saveResults(db db.Manager, resCh chan OrderBook, collection string) {
	for r := range resCh {
		err := db.Write(r, collection)
		if err != nil {
			fmt.Println("Error during write to DB", err)
		}
	}
}