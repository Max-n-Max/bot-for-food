package ticker

import (
	"fmt"
	"github.com/Max-n-Max/bot-for-food/exchange"
	"time"
)

type Manager struct {
	ExchangeManager exchange.Manager
	orderBookCh     chan string
	tradesCh        chan string
	orderBookTicker *time.Ticker
	tradesTicker    *time.Ticker
}

func NewManager(manager exchange.Manager, orderBookResCh chan string, tradesResCh chan string) *Manager {
	c := new(Manager)
	c.ExchangeManager = manager
	c.orderBookCh = orderBookResCh
	c.tradesCh = tradesResCh
	// Get tickers from config
	c.orderBookTicker = time.NewTicker(10000 * time.Millisecond)
	c.tradesTicker = time.NewTicker(10000 * time.Millisecond)

	return c
}

func (c *Manager) Run() {
	go c.runOrderBookTicker()
	go c.runTradesTicker()
}

func (c *Manager) runOrderBookTicker() {
	c.runTicker(*c.orderBookTicker, c.orderBookCh)
}

func (c *Manager) runTradesTicker() {
	c.runTicker(*c.tradesTicker, c.tradesCh)
}

func (c Manager) runTicker(ticker time.Ticker, channel chan string) {
	for {
		_ := <-ticker.C
		res, e := c.ExchangeManager.GetOrderBook()
		if e == nil {
			fmt.Println("Got data from exchange: ", res)
			channel <- res
		} else {
			fmt.Println("Cannot get data from exchange. Error: ", e)
		}
	}
}

