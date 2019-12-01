package http

import (
	"github.com/Max-n-Max/bot-for-food/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Manager struct {
	handler Handler
	config  config.Manager
}

func NewManager(handler Handler, config config.Manager) *Manager {
	m := new(Manager)
	m.handler = handler
	m.config = config

	return m
}

func (m *Manager) Run() {
	log.Println("Starting server...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/get_bot_info",            m.handler.GetBotInfoHandler)       .Methods("POST")
	router.HandleFunc("/get_order_book",          m.handler.GetOrderBookHandler)     .Methods("POST")
	router.HandleFunc("/get_candles_history",     m.handler.GetCandlesHistoryHandler).Methods("POST")
	router.HandleFunc("/collect/orderbook/start", m.handler.StartCollectorHandler)   .Methods("POST")
	router.HandleFunc("/collect/orderbook/stop",  m.handler.StopCollectorHandler)    .Methods("POST")

	// for static
	// http://localhost:9090/app/
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/client_src/"))).Methods("GET")

	address := ":" + "9090"
	log.Fatal(http.ListenAndServe(address, router))
}

