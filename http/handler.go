package http

import (
	"encoding/json"
	"fmt"
	"github.com/Max-n-Max/bot-for-food/config"
	"github.com/Max-n-Max/bot-for-food/db"
	"github.com/Max-n-Max/bot-for-food/exchange"
	"github.com/Max-n-Max/bot-for-food/resources"
	"net/http"
)

type Handler struct {
	db       *db.Manager
	exchange exchange.Manager
	config   config.Manager
}


func NewHandler(db *db.Manager, exchange exchange.Manager, config config.Manager) *Handler{
	h := new(Handler)
	h.db = db
	h.exchange = exchange
	h.config = config

	return h
}

func (h *Handler) OrderBookHandler(w http.ResponseWriter, r *http.Request) {
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
	w.Write([]byte(res))
	w.Header().Set("Content-Type", "application/json")
}

func (h *Handler) CandlesHistoryHandler(w http.ResponseWriter, r *http.Request) {
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