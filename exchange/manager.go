package exchange

import (
	"encoding/json"
	"github.com/bitfinexcom/bitfinex-api-go/v1"
)

type Manager struct {
	client* bitfinex.Client
}

func NewManager() *Manager {
	m := new(Manager)
	m.client = bitfinex.NewClient()
	m.client.Auth("22dsdssddssd", "2dssdsdsdsdsd22")
	return m
}

func (m *Manager) GetOrderBook() (string, error) {
	orders, err := m.client.OrderBook.Get(bitfinex.BTCUSD, 10, 10, false)
	if err != nil {
		return "", err
	}
	out, err := json.Marshal(orders)
	if err != nil {
		return "", err
	}

	return string(out), nil
}


func (m *Manager) GetTrades() (string, error) {
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
