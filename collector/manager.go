package collector

import (
	"context"
	"fmt"
	"github.com/Max-n-Max/bot-for-food/config"
	"github.com/Max-n-Max/bot-for-food/db"
	"github.com/Max-n-Max/bot-for-food/exchange"
	"github.com/Max-n-Max/bot-for-food/resources"
	"github.com/bitfinexcom/bitfinex-api-go/v2/websocket"
	"log"
)

type orderBookCollector struct {
	in  chan string
	out chan resources.OrderBook
	c   *websocket.Client
	id  string
}

func (o *orderBookCollector) unsubscribe() {
	o.c.Unsubscribe(context.Background(), o.id)
	close(o.in)
	close(o.out)
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

func (m *Manager) StartOrderBookCollection(pair, precision, frequency string, priceLevel int) {
	log.Println("Start", pair, "collecting")

	// register listener with given params
	pc := orderBookCollector{in: make(chan string), out: make(chan resources.OrderBook), c: websocket.New()}
	m.pool[pair] = &pc

	id, err := startOrderBookListener(m.pool[pair].c,
		m.ExchangeManager.GetOrderBook,
		pair,
		precision,
		frequency,
		priceLevel,
		m.pool[pair].in)
	if err != nil {
		return
	}

	m.pool[pair].id = id

	go ProcessOrderBook(m.pool[pair].in, m.pool[pair].out)
	go saveResults(m.db, m.pool[pair].out, m.config.GetString("db.order-book-collection"))
}

func (m *Manager) StopCollection(pair string) {
	log.Println("Stop", pair, "collecting")
	m.pool[pair].unsubscribe()
	delete(m.pool, pair)
}

func saveResults(db db.Manager, resCh chan resources.OrderBook, collection string) {
	for r := range resCh {
		err := db.Write(r, collection)
		if err != nil {
			fmt.Println("Error during write to DB", err)
		}
	}
}
