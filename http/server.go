package http

import (
	"encoding/json"
	"fmt"
	"github.com/Max-n-Max/bot-for-food/config"
	"github.com/Max-n-Max/bot-for-food/db"
	"github.com/Max-n-Max/bot-for-food/exchange"
	"github.com/Max-n-Max/bot-for-food/resources"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Manager struct {
	db       *db.Manager
	exchange exchange.Manager
	config   config.Manager
}

func NewManager(db *db.Manager, exchange exchange.Manager, config config.Manager) *Manager {
	m := new(Manager)
	m.db = db
	m.exchange = exchange
	m.config = config

	return m
}

func (m *Manager) Run() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/orderbook", m.orderBookHandler).Methods("POST")
	router.HandleFunc("/candles/history", m.candlesHistoryHandler).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/client_src/"))).Methods("GET") //http://localhost:9090/app/

	address := ":" + "9090"
	log.Fatal(http.ListenAndServe(address, router))

}

func (m *Manager) orderBookHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rBody resources.OrderBookReqBody
	err := decoder.Decode(&rBody)
	if err != nil {
		panic(err)
	}

	res, err := m.db.QueryOrderBook(rBody.From, rBody.To, rBody.Pair, m.config.GetString("db.order-book-collection"))
	if err != nil {
		fmt.Println("ERROR")
		//TODO return ERROR
	}
	w.Write([]byte(res))
	w.Header().Set("Content-Type", "application/json")
}

func (m *Manager) candlesHistoryHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rBody resources.CandlesHistoryBody
	err := decoder.Decode(&rBody)
	if err != nil {
		panic(err)
	}

	res, err := m.exchange.GetCandlesHistory(rBody)

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
