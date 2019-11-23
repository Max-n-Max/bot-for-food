package collector

import (
	"fmt"
	"github.com/Max-n-Max/bot-for-food/exchange"
	"log"
	"time"
)

type Manager struct {
	ExchangeManager exchange.Manager
	orderBookCh     chan string
	tradesCh        chan string

}

func NewManager(manager exchange.Manager) *Manager{
	log.Println("Starting collector...")

	c := new(Manager)
	c.ExchangeManager = manager

	return c
}

type queryType func(pair string) (string, error)

func (c *Manager) RunOrderBookTicker(pair string, interval int, orderBookResCh chan string) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	runTicker(c.ExchangeManager.GetOrderBook, pair, *ticker, orderBookResCh)
}

func (c *Manager) RunTradesTicker(pair string, interval int,tradesResCh chan string) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	runTicker(c.ExchangeManager.GetTrades, pair, *ticker, tradesResCh)
}

func runTicker(fn queryType, pair string, ticker time.Ticker, channel chan string) {
	for {
		_ = <-ticker.C
		res, e := fn(pair)
		if e == nil {
			//fmt.Println("Got data from exchange: ", res)
			channel <- res
		} else {
			fmt.Println("Cannot get data from exchange. Error: ", e)
		}
	}
}

