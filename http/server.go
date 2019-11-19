package http

import (
	"encoding/json"
	"fmt"
	"github.com/Max-n-Max/bot-for-food/config"
	"github.com/Max-n-Max/bot-for-food/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Manager struct {
	db     *db.Manager
	config config.Manager
}

func NewManager(db *db.Manager, config config.Manager) *Manager {
	m := new(Manager)
	m.db = db
	m.config = config

	return m
}

func (m *Manager) Run() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/get_order_book", m.dataGetHandler).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/client_src/"))).Methods("GET") //http://localhost:9090/app/

	address := ":" + "9090"
	log.Fatal(http.ListenAndServe(address, router))

}

type OrderBookReqBody struct {
	From string `json:"date_start"`
	To   string `json:"date_end"`
	Pair string `json:"pair"`
}

func (m *Manager) dataGetHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rBody OrderBookReqBody
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
