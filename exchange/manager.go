package exchange

import (
	"encoding/json"
	"github.com/Max-n-Max/bot-for-food/config"
	"github.com/Max-n-Max/bot-for-food/resources"
	"github.com/bitfinexcom/bitfinex-api-go/v1"
	bitfinexV2 "github.com/bitfinexcom/bitfinex-api-go/v2"
	v2 "github.com/bitfinexcom/bitfinex-api-go/v2/rest"
	"log"
)

type Manager struct {
	clientV1 *bitfinex.Client
	clientV2 *v2.Client

	config config.Manager
}

func NewManager(config config.Manager) *Manager {
	log.Println("Starting exchange...")
	m := new(Manager)
	m.clientV1 = bitfinex.NewClient()
	m.clientV1.Auth(config.GetString("key"), config.GetString("secret"))
	m.clientV2 = v2.NewClient()
	m.config = config
	return m
}

func (m *Manager) GetOrderBook(pair string) (string, error) {
	orders, err := m.clientV1.OrderBook.Get(
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

func (m *Manager) GetCandlesHistory(body resources.CandlesHistoryBody) (*bitfinexV2.CandleSnapshot, error) {

	var result *bitfinexV2.CandleSnapshot
	var err error

	if body.Start != 0 && body.End != 0 && body.Limit != 0 {
		var sort int
		if body.OldestFirst {
			sort = 1
		} else {
			sort = -1
		}

		result, err = m.clientV2.Candles.HistoryWithQuery(
			bitfinexV2.TradingPrefix+body.Pair,
			bitfinexV2.CandleResolution(body.Resolution),
			bitfinexV2.Mts(body.Start),
			bitfinexV2.Mts(body.End),
			bitfinexV2.QueryLimit(body.Limit),
			bitfinexV2.SortOrder(sort),
		)
	} else {
		result, err = m.clientV2.Candles.History(
			bitfinexV2.TradingPrefix+body.Pair,
			bitfinexV2.CandleResolution(body.Resolution),
		)
	}

	return result, err

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
