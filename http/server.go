package http

import (
	"github.com/Max-n-Max/bot-for-food/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Manager struct {
	handler Handler
	config   config.Manager
}

func NewManager(handler Handler, config config.Manager) *Manager {
	m := new(Manager)
	m.handler = handler
	m.config = config

	return m
}

func (m *Manager) Run() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/get_order_book", m.dataGetHandler).Methods("POST")
	router.HandleFunc("/candles/history", m.handler.CandlesHistoryHandler).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/client_src/"))).Methods("GET") //http://localhost:9090/app/

	address := ":" + "9090"
	log.Fatal(http.ListenAndServe(address, router))
}

type OrderBookReqBody struct {
	From string `json:"date_start"`
	To   string `json:"date_end"`
	Pair string `json:"pair"`
}

