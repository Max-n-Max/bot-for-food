package main

import (
	"encoding/json"
	"fmt"
	"github.com/Max-n-Max/bot-for-food/exchange"
	"github.com/Max-n-Max/bot-for-food/ticker"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

const (
	hosts      = "localhost:27017"
	database   = "cryptodb"
	username   = ""
	password   = ""
	collection = "trades"
)

type Book struct {
	Price     string `json:"Price"`
	Rate      string `json:"Rate"`
	Amount    string `json:"Amount"`
	Period    int    `json:"Period"`
	Timestamp string `json:"Timestamp"`
	Frr       string `json:"Frr"`
}

type Job struct {
	Bids []Book
	Asks []Book
}

type MongoStore struct {
	session *mgo.Session
}

var mongoStore = MongoStore{}

func main() {

	session := initialiseMongo()
	mongoStore.session = session

	exchangeManager := exchange.NewManager()

	resultChan := make(chan string)
	stopChan := make(chan bool)
	// Create and init Manager
	collector := ticker.NewManager(*exchangeManager, resultChan, stopChan)
	collector.Run()

	go func() {
		for r := range resultChan {
			write(r)
		}
	}()

	//Create MongoDB session

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/data", dataGetHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":9090", router))

}

func initialiseMongo() (session *mgo.Session) {

	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}
	return

}

func dataGetHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	col := mongoStore.session.DB(database).C(collection)

	results := []Job{}
	col.Find(bson.M{"title": bson.RegEx{"", ""}}).All(&results)
	jsonString, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(jsonString))

}

func write(record string) {
	col := mongoStore.session.DB(database).C(collection)
	fmt.Println("Going to insert to BD", record)
	//Save data into Job struct
	var job Job
	err := json.Unmarshal([]byte(record), &job)
	if err != nil {
		fmt.Println("Error during parsing results", err)
		return
	}

	//Insert job into MongoDB
	err = col.Insert(job)
	if err != nil {
		fmt.Println("Error during insert to DB", err)
	}
}

//TODO Graceful shutdown
