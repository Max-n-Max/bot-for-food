package resources

import "github.com/bitfinexcom/bitfinex-api-go/v1"

type OrderBook struct {
	Timestamp string
	Pair      string
	BidsWall  float64
	AsksWall  float64
	Window    float64
	Bids      []bitfinex.OrderBookEntry
	Asks      []bitfinex.OrderBookEntry
}


type OrderBookReqBody struct {
	From string `json:"date_start"`
	To   string `json:"date_end"`
	Pair string `json:"pair"`
}