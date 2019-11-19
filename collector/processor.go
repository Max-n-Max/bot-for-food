package collector

import (
	"encoding/json"
	"fmt"
	"github.com/Max-n-Max/bot-for-food/resources"
	"github.com/bitfinexcom/bitfinex-api-go/v1"
	"strconv"
	"time"
)

func ProcessOrderBook(in chan string, out chan resources.OrderBook) {
	for j := range in {
		var job resources.OrderBook
		err := json.Unmarshal([]byte(j), &job)
		if err != nil {
			fmt.Println("Error during parsing results", err)
			return
		}
		job.Timestamp = time.Now().String()
		job.BidsWall = getWall(job.Bids)
		job.AsksWall = getWall(job.Asks)
		job.Window = 2 * (job.AsksWall - job.BidsWall) / (job.AsksWall + job.BidsWall)

		out <- job
	}
}

func getWall(asks []bitfinex.OrderBookEntry) float64 {
	for _, o := range asks {
		price, err := strconv.ParseFloat(o.Price, 64)
		if err != nil {
			fmt.Println("Cannot convert price: ", o.Price, err)
		}
		amount, err := strconv.ParseFloat(o.Amount, 64)
		if err != nil {
			fmt.Println("Cannot convert amount: ", o.Amount, err)
		}

		if price*amount >= 1000 {
			return price
		}
	}

	lastOrder := asks[len(asks)-1]
	lastPrice, err := strconv.ParseFloat(lastOrder.Price, 64)
	if err != nil {
		fmt.Println("Cannot convert price: ", lastOrder.Price, err)
	}
	return lastPrice
}
