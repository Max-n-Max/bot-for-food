package main

import (
	"github.com/Max-n-Max/bot-for-food/cmd"
	"github.com/Max-n-Max/bot-for-food/collector"
	"github.com/Max-n-Max/bot-for-food/db"
	"github.com/Max-n-Max/bot-for-food/exchange"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)



func main() {

	exchange := exchange.NewManager()
	db, err := db.NewManager()
	if err != nil {
		panic(err)
	}
	collector := collector.NewManager(*exchange)


	cmd.Execute(*collector, *db)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/data", dataGetHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":9090", router))

}


func dataGetHandler(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//
	//col := mongoStore.session.DB(database).C(collection)
	//
	//results := []Job{}
	//col.Find(bson.M{"title": bson.RegEx{"", ""}}).All(&results)
	//jsonString, err := json.Marshal(results)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Fprint(w, string(jsonString))

}


//TODO Graceful shutdown
