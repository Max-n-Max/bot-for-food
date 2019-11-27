package collector

import (
	"fmt"
	"github.com/Max-n-Max/bot-for-food/config"
	"github.com/Max-n-Max/bot-for-food/db"
	"github.com/Max-n-Max/bot-for-food/exchange"
	"github.com/Max-n-Max/bot-for-food/resources"
	"log"
	"time"
)

type orderBookCollector struct {
	in  chan string
	out chan resources.OrderBook
	t   *time.Ticker
}

type Manager struct {
	ExchangeManager exchange.Manager
	db              db.Manager
	config          config.Manager
	pool            map[string]*orderBookCollector
	orderBookCh     chan string
	tradesCh        chan string
}

type queryType func(pair string) (string, error)

func NewManager(manager exchange.Manager, db db.Manager, config config.Manager) *Manager {
	log.Println("Starting collector...")

	m := new(Manager)
	m.db = db
	m.ExchangeManager = manager
	m.pool = make(map[string]*orderBookCollector)
	m.config = config

	return m
}

func (m *Manager) StartOrderBookCollection(pair string, interval int) {
	log.Println("Start ", pair, "collecting")
	pc := orderBookCollector{in: make(chan string), out: make(chan resources.OrderBook), t: time.NewTicker(time.Duration(interval) * time.Second)}
	m.pool[pair] = &pc
	go runTicker(m.ExchangeManager.GetOrderBook, pair, *m.pool[pair].t, m.pool[pair].in)
	go ProcessOrderBook(m.pool[pair].in, m.pool[pair].out)
	go saveResults(m.db, m.pool[pair].out, m.config.GetString("db.order-book-collection"))
}

func (m *Manager) StopCollection(pair string) {
	log.Println("Stop ", pair, "collecting")
	m.pool[pair].t.Stop()
	close(m.pool[pair].in)
	close(m.pool[pair].out)
	delete(m.pool, pair)
}

func runTicker(fn queryType, pair string, ticker time.Ticker, channel chan string) {
	for {
		_ = <-ticker.C
		res, e := fn(pair)
		if e == nil {
			channel <- res
		} else {
			fmt.Println("Cannot get data from exchange. Error: ", e)
		}
	}
}

func saveResults(db db.Manager, resCh chan resources.OrderBook, collection string) {
	for r := range resCh {
		err := db.Write(r, collection)
		if err != nil {
			fmt.Println("Error during write to DB", err)
		}
	}
}
