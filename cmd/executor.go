package cmd

import (
	"fmt"
	"github.com/Max-n-Max/bot-for-food/collector"
	"github.com/Max-n-Max/bot-for-food/config"
	"github.com/Max-n-Max/bot-for-food/db"
	"github.com/Max-n-Max/bot-for-food/resources"
)


func Execute(collector collector.Manager, db db.Manager, config config.Manager) {

	orderBookExchangeCh := make(chan string)
	//tradesExchangeCh := make(chan string)

	orderBookProcessCh := make(chan resources.OrderBook)
	//tradesProcessCh := make(chan OrderBook)
	pairs := config.GetStringSlice("exchange.pairs")
	for _, pair := range pairs {
		go collector.RunOrderBookTicker(pair, config.GetInt("order-book-ticker"), orderBookExchangeCh)
	}
	//go collector.RunTradesTicker(config.GetInt("trades-ticker"), tradesExchangeCh)

	go ProcessOrderBook(orderBookExchangeCh, orderBookProcessCh)
	go saveResults(db, orderBookProcessCh, config.GetString("db.order-book-collection"))
}

func saveResults(db db.Manager, resCh chan resources.OrderBook, collection string) {
	for r := range resCh {
		err := db.Write(r, collection)
		if err != nil {
			fmt.Println("Error during write to DB", err)
		}
	}
}
