package http

import (
	"context"
	"flag"
	"github.com/Max-n-Max/bot-for-food/config"
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	websocetBF "github.com/bitfinexcom/bitfinex-api-go/v2/websocket"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
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

	finish := make(chan bool)

	log.Println("Starting server...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/get_order_book",      m.handler.OrderBookHandler     ).Methods("POST")
	router.HandleFunc("/get_candles_history", m.handler.CandlesHistoryHandler).Methods("POST")

	// for static
	// http://localhost:9090/app/
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/client_src/"))).Methods("GET")


	go func() {
		address := ":" + "9090"
		log.Fatal(http.ListenAndServe(address, router))
	}()

	routerWs := mux.NewRouter()
	routerWs.HandleFunc("/ws", serveWs)
	go func() {
		addressWs := ":" + "9092"
		log.Fatal(http.ListenAndServe(addressWs, routerWs))
	}()

	<-finish
}


const (
	// Time allowed to write the file to the client.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = 60 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Poll file for changes with this period.
	filePeriod = 10 * time.Second
)

var (
	addr      = flag.String("addr", ":8092", "http service address")
	filename  = "/Users/maxim/Appsflyer/projects/DEMOS/BranchIo/ios-branch-deep-linking-attribution/Branch-SDK/Branch-SDK/BNCAppleReceipt.h"
	upgrader  = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func readFileIfModified(lastMod time.Time) ([]byte, time.Time, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return nil, lastMod, err
	}
	if !fi.ModTime().After(lastMod) {
		return nil, lastMod, nil
	}
	p, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fi.ModTime(), err
	}
	return p, fi.ModTime(), nil
}

func reader(ws *websocket.Conn) {
	defer ws.Close()
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}

func writer(ws *websocket.Conn, lastMod time.Time) {
	lastError := ""
	pingTicker := time.NewTicker(pingPeriod)
	fileTicker := time.NewTicker(filePeriod)
	defer func() {
		pingTicker.Stop()
		fileTicker.Stop()
		ws.Close()
	}()
	for {
		select {
		case <-fileTicker.C:
			var p []byte
			var err error

			p, lastMod, err = readFileIfModified(lastMod)

			if err != nil {
				if s := err.Error(); s != lastError {
					lastError = s
					p = []byte(lastError)
				}
			} else {
				lastError = ""
			}

			if p != nil {
				ws.SetWriteDeadline(time.Now().Add(writeWait))
				if err := ws.WriteMessage(websocket.TextMessage, p); err != nil {
					return
				}
			}
		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	} // to get rid of websocket: request origin not allowed by Upgrader.CheckOrigin

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	client := websocetBF.New()
	err = client.Connect()
	if err != nil {
		log.Printf("could not connect: %s", err.Error())
		return
	}

	_, err = client.SubscribeTrades(context.Background(), bitfinex.TradingPrefix + "ETPUSD")
	if err != nil {
		log.Fatal(err)
	}

	for obj := range client.Listen() {
		switch obj.(type) {
		case error:
			log.Printf("EROR RECV: %s", obj)
		default:
			log.Printf("MSG RECV: %#v", obj)
			//p := []byte(obj)
			//
			//ws.SetWriteDeadline(time.Now().Add(writeWait))
			//if err := ws.WriteMessage(websocket.TextMessage, p); err != nil {
			//	return
			//}
		}
	}





	//var lastMod time.Time
	//if n, err := strconv.ParseInt(r.FormValue("lastMod"), 16, 64); err == nil {
	//	lastMod = time.Unix(0, n)
	//}
	//fmt.Println("lastMod: ", lastMod)
	//go writer(ws, lastMod)
	//reader(ws)
}

