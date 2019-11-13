package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func ProcessOrderBook(in chan string, out chan OrderBook) {
	for j := range in {
		var job OrderBook
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

func getWall(asks []Order) float64 {
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
