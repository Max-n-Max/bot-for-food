package collector

import (
	"fmt"
	"github.com/Max-n-Max/bot-for-food/exchange"
	"time"
)

type Manager struct {
	ExchangeManager exchange.Manager
	orderBookCh     chan string
	tradesCh        chan string

}

func NewManager(manager exchange.Manager) *Manager{
	c := new(Manager)
	c.ExchangeManager = manager

	return c
}

type queryType func() (string, error)

func (c *Manager) RunOrderBookTicker(orderBookResCh chan string) {
	// TODO get tickers from cong
	ticker := time.NewTicker(10000 * time.Millisecond)
	runTicker(c.ExchangeManager.GetOrderBook, *ticker, orderBookResCh)
}

func (c *Manager) RunTradesTicker(tradesResCh chan string) {
	// TODO get tickers from cong
	ticker := time.NewTicker(10000 * time.Millisecond)
	runTicker(c.ExchangeManager.GetTrades, *ticker, tradesResCh)
}

func runTicker(fn queryType,ticker time.Ticker, channel chan string) {
	for {
		_ = <-ticker.C
		res, e := fn()
		if e == nil {
			fmt.Println("Got data from exchange: ", res)
			channel <- res
		} else {
			fmt.Println("Cannot get data from exchange. Error: ", e)
		}
	}
}

