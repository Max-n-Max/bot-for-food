package collector

import (
	"encoding/json"
	"fmt"
	"github.com/Max-n-Max/bot-for-food/resources"
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
		out <- job
	}
}
