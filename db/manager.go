package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
)

type Manager struct {
	session *mgo.Session
}

func NewManager() (*Manager, error){
	m := new(Manager)

	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}
	session, err := mgo.DialWithInfo(info)
	m.session = session
	return m, err
}

const (
	hosts      = "localhost:27017"
	database   = "cryptodb"
	username   = ""
	password   = ""
	tradesCollection = "trades"
	orderBookCollection = "orderbook"
)


func (m *Manager) Write(record interface{}, collection string) error{
	col := m.session.DB(database).C(collection)
	fmt.Println("Going to insert to BD", record)


	//Insert job into MongoDB
	err := col.Insert(record)
	if err != nil {
		fmt.Println("Error during insert to DB", err)
	}

	return err
}


func (m *Manager) Query(query string) (string, error) {
	// TODO query DB

	return "", nil
}