package exchange

import (
	"encoding/json"
	"github.com/Max-n-Max/bot-for-food/config"
	"github.com/Max-n-Max/bot-for-food/resources"
	"github.com/bitfinexcom/bitfinex-api-go/v1"
)

type Manager struct {
	client *bitfinex.Client
	config config.Manager
}

func NewManager(config config.Manager) *Manager {
	m := new(Manager)
	m.client = bitfinex.NewClient()
	m.client.Auth(config.GetString("key"), config.GetString("secret"))
	m.config = config
	return m
}

func (m *Manager) GetOrderBook(pair string) (string, error) {
	orders, err := m.client.OrderBook.Get(
		pair,
		m.config.GetInt("exchange.bids-limit"),
		m.config.GetInt("exchange.asks-limit"),
		m.config.GetBool("exchange.no-group"))
	if err != nil {
		return "", err
	}

	var job = resources.OrderBook{Bids: orders.Bids, Asks: orders.Asks, Pair: pair}
	out, err := json.Marshal(job)
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func (m *Manager) GetTrades(pair string) (string, error) {
	//orders, err := m.client.OrderBook.Get(bitfinex.BTCUSD, 10, 10, false)
	//if err != nil {
	//	return "", err
	//}
	//out, err := json.Marshal(orders)
	//if err != nil {
	//	return "", err
	//}

	return string(""), nil
}
