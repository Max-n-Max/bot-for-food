package resources

import "github.com/bitfinexcom/bitfinex-api-go/v1"

type OrderBook struct {
	Timestamp string
	Pair      string
	//BidsWall  float64
	//AsksWall  float64
	//Window    float64
	Bids []bitfinex.OrderBookEntry
	Asks []bitfinex.OrderBookEntry
}

type OrderBookRes struct {
	Timestamp string
	Pair      string
	BidsWall  float64
	AsksWall  float64
	Window    float64
	Bids      []bitfinex.OrderBookEntry
	Asks      []bitfinex.OrderBookEntry
}

type OrderBookReqBody struct {
	From    string  `json:"from"`
	To      string  `json:"to"`
	Pair    string  `json:"pair"`
	Wall    float64 `json:"wall"`
	SumWall float64 `json:"sumwall"`
	Window  float64 `json:"window"`
}
