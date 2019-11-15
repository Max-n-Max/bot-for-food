package http

import (
	"fmt"
	"github.com/Max-n-Max/bot-for-food/config"
	"github.com/Max-n-Max/bot-for-food/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Manager struct {
	db *db.Manager
	config config.Manager
}

func NewManager(db *db.Manager, config config.Manager) *Manager{
	m := new(Manager)
	m.db = db
	m.config = config

	return m
}

func (m *Manager) Run() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/orderbook", m.dataGetHandler).Methods("POST")
	address := ":" + "9090"
	log.Fatal(http.ListenAndServe(address, router))
}

func (m *Manager) dataGetHandler(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	res ,err := m.db.QueryOrderBook(from, to, m.config.GetString("db.order-book-collection"))
	if err != nil {
		fmt.Println("ERROR")
		//TODO return ERROR
	}
	w.Write([]byte(res))
	w.Header().Set("Content-Type", "application/json")


}