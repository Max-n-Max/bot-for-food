package db

import (
	"encoding/json"
	"fmt"
	"github.com/Max-n-Max/bot-for-food/config"
	"github.com/Max-n-Max/bot-for-food/resources"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type Manager struct {
	session *mgo.Session
	dbName string
}

func NewManager(config config.Manager) (*Manager, error){
	log.Println("Starting DB...")
	m := new(Manager)
	m.dbName = config.GetString("db.database")
	info := &mgo.DialInfo{
		Addrs:    []string{config.GetString("db.hosts")},
		Timeout:  time.Duration(config.GetInt("db.timeout")) * time.Second,
		Database: config.GetString("db.database"),
		Username: config.GetString("db.username"),
		Password: config.GetString("db.password"),
	}
	session, err := mgo.DialWithInfo(info)
	m.session = session
	return m, err
}


func (m *Manager) Write(record interface{}, collection string) error{
	col := m.session.DB(m.dbName).C(collection)
	//fmt.Println("Going to insert to BD", record)


	//Insert job into MongoDB
	err := col.Insert(record)
	if err != nil {
		fmt.Println("Error during insert to DB", err)
	}

	return err
}


func (m *Manager) QueryOrderBook(from, to, pair, collection string) (string, error) {
	var results []resources.OrderBook

	//r := record{Timestamp:timestamp{gte:"2019-11-13", lt:"2019-11-14"}}
	col := m.session.DB(m.dbName).C(collection)
	_ = col.Find(bson.M{"timestamp": bson.M{"$gt": from, "$lt": to}, "pair":pair}).All(&results)
	b, err := json.Marshal(results)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(b), nil
}