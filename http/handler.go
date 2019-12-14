package http

import (
	"encoding/json"
	"fmt"
	"github.com/Max-n-Max/bot-for-food/collector"
	"github.com/Max-n-Max/bot-for-food/config"
	"github.com/Max-n-Max/bot-for-food/db"
	"github.com/Max-n-Max/bot-for-food/exchange"
	"github.com/Max-n-Max/bot-for-food/resources"
	"github.com/bitfinexcom/bitfinex-api-go/v1"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	db        *db.Manager
	exchange  exchange.Manager
	collector *collector.Manager
	config    config.Manager
}

func NewHandler(collector *collector.Manager, db *db.Manager, exchange exchange.Manager, config config.Manager) *Handler {
	log.Println("Starting http handler...")
	h := new(Handler)
	h.collector = collector
	h.db = db
	h.exchange = exchange
	h.config = config

	return h
}

func (h *Handler) StartCollectorHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rBody resources.StartOrderBookCollectorBody
	err := decoder.Decode(&rBody)
	if err != nil {
		panic(err)
	}

	h.collector.StartOrderBookCollection(rBody.Pair, rBody.Precision, rBody.Frequency, rBody.PriceLevel)
}

func (h *Handler) StopCollectorHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rBody resources.StopOrderBookCollectorBody
	err := decoder.Decode(&rBody)
	if err != nil {
		panic(err)
	}

	h.collector.StopCollection(rBody.Pair)
}

func (h *Handler) GetBotInfoHandler(w http.ResponseWriter, r *http.Request){

}

func (h *Handler) GetOrderBookHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rBody resources.OrderBookReqBody
	err := decoder.Decode(&rBody)
	if err != nil {
		panic(err)
	}

	res, err := h.db.QueryOrderBook(rBody.From, rBody.To, rBody.Pair, h.config.GetString("db.order-book-collection"))
	if err != nil {
		fmt.Println("ERROR")
		//TODO return ERROR
	}

	// Calculate walls and window and add it to result
	enOB := enrichOrderBook(res, rBody.Wall, rBody.SumWall, rBody.Window, rBody.SkipB)
	responseBody := enOB

	if responseBody == "" || responseBody == "null"{
		responseBody = "{}"
	}
	w.Write([]byte(responseBody))
	w.Header().Set("Content-Type", "application/json")
}

func (h *Handler) GetCandlesHistoryHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rBody resources.CandlesHistoryBody
	err := decoder.Decode(&rBody)
	if err != nil {
		panic(err)
	}

	res, err := h.exchange.GetCandlesHistory(rBody)

	if err != nil {
		fmt.Println("ERROR")
		//TODO return ERROR
		http.Error(w, "Error!!!", http.StatusBadRequest)
	}
	resJson, err := json.Marshal(res)
	if err != nil {
		fmt.Println("ERROR")
		//TODO return ERROR
		http.Error(w, "Error!!!", http.StatusBadRequest)
	} else {
		w.Write([]byte(resJson))
		w.Header().Set("Content-Type", "application/json")
	}

}

//TODO: worth to move to separate module, (in current handler we deal with http only)
func enrichOrderBook(orderbook string, wall, sumWall float64, reqWindow float64, skipOrderBookData bool) string {
	var OB []resources.OrderBook
	var res []resources.OrderBookRes
	err := json.Unmarshal([]byte(orderbook), &OB)
	if err != nil {
		panic(err)
	}
	for _, o := range OB {
		aWall := getWall(o.Asks, wall, sumWall)
		bWall := getWall(o.Bids, wall, sumWall)
		var window = 0.0
		if aWall != 0 && bWall != 0 {
			window = 2 * (aWall - bWall) / (aWall + bWall)
		}

		if window >= reqWindow {

			var bids []bitfinex.OrderBookEntry
			var asks []bitfinex.OrderBookEntry
			if !skipOrderBookData {
				bids = o.Bids
				asks = o.Asks
			}

			copyOrder := resources.OrderBookRes{
				Timestamp: o.Timestamp,
				Pair:      o.Pair,
				Bids:      bids,
				Asks:      asks,
				BidsWall:  bWall,
				AsksWall:  aWall,
				Window:    window,
			}

			res = append(res, copyOrder)
		}
	}
	resJson, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	return string(resJson)
}

func getWall(asks []bitfinex.OrderBookEntry, wall, sumWall float64) float64 {
	var sum float64 = 0
	for _, o := range asks {
		price, err := strconv.ParseFloat(o.Price, 64)
		if err != nil {
			fmt.Println("Cannot convert price: ", o.Price, err)
		}
		amount, err := strconv.ParseFloat(o.Amount, 64)
		if err != nil {
			fmt.Println("Cannot convert amount: ", o.Amount, err)
		}
		sum += price * amount

		if sum >= sumWall && price*amount >= wall {
			return price
		}
	}

	return 0
}
