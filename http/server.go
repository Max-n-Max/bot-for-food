package http

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RunServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/data", dataGetHandler).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/client_src/"))).Methods("GET") //http://localhost:9090/app/
	address := ":" + "9090"
	log.Fatal(http.ListenAndServe(address, router))

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