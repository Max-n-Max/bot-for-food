package main

import (
	"github.com/Max-n-Max/bot-for-food/cmd"
	"github.com/Max-n-Max/bot-for-food/collector"
	"github.com/Max-n-Max/bot-for-food/config"
	"github.com/Max-n-Max/bot-for-food/db"
	"github.com/Max-n-Max/bot-for-food/exchange"
	"github.com/Max-n-Max/bot-for-food/http"
)



func main() {
	config := config.NewManager("./config", "./config/creds", "config")
	exchange := exchange.NewManager(config)
	db, err := db.NewManager(config)
	if err != nil {
		panic(err)
	}
	collector := collector.NewManager(*exchange)

	cmd.Execute(*collector, *db, config)
	http.RunServer()
}

//TODO Graceful shutdown
