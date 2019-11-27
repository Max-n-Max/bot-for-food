package http

import (
	"encoding/json"
	_ "fmt"
	"github.com/Max-n-Max/bot-for-food/resources"
	"github.com/bitfinexcom/bitfinex-api-go/v1"
	"github.com/stretchr/testify/assert"
	_ "io/ioutil"
	_ "os"
	"testing"
)

func TestEnrich(t *testing.T) {
	byteValue, _ := json.Marshal(createTestDataset())

	res := enrichOrderBook(string(byteValue), 2000.0, 1000.0, 0.01)
	assert.NotNil(t, res, "")
}

func createTestDataset() []resources.OrderBook {
	var res []resources.OrderBook

	e1Bid1 := bitfinex.OrderBookEntry{Price:"999", Amount:"0.55"}
	e1Bid2 := bitfinex.OrderBookEntry{Price:"998", Amount:"0.55"}
	e1Bid3 := bitfinex.OrderBookEntry{Price:"990", Amount:"3"}

	e1Ask1 := bitfinex.OrderBookEntry{Price:"1000", Amount:"0.5"}
	e1Ask2 := bitfinex.OrderBookEntry{Price:"1001", Amount:"0.5"}
	e1Ask3 := bitfinex.OrderBookEntry{Price:"1002", Amount:"2"}

	var e1Bids []bitfinex.OrderBookEntry
	var e1Asks []bitfinex.OrderBookEntry

	e1Bids = append(e1Bids, e1Bid1, e1Bid2, e1Bid3)
	e1Asks = append(e1Asks, e1Ask1, e1Ask2, e1Ask3)


	e1 := resources.OrderBook{Bids: e1Bids, Asks:e1Asks}
	res = append(res, e1)

	return res

}
