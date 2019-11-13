package main

import (
	"github.com/Max-n-Max/bot-for-food/http"
	"github.com/Max-n-Max/bot-for-food/cmd"
	"github.com/Max-n-Max/bot-for-food/collector"
	"github.com/Max-n-Max/bot-for-food/db"
	"github.com/Max-n-Max/bot-for-food/exchange"
)



func main() {

	exchange := exchange.NewManager()
	db, err := db.NewManager()
	if err != nil {
		panic(err)
	}
	collector := collector.NewManager(*exchange)


	cmd.Execute(*collector, *db)
	http.RunServer()
}

//TODO Graceful shutdown
