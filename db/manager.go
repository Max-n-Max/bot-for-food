package db

import (
	"fmt"
	"github.com/Max-n-Max/bot-for-food/config"
	"gopkg.in/mgo.v2"
	"time"
)

type Manager struct {
	session *mgo.Session
	dbName string
}

func NewManager(config config.Manager) (*Manager, error){
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
	fmt.Println("Going to insert to BD", record)


	//Insert job into MongoDB
	err := col.Insert(record)
	if err != nil {
		fmt.Println("Error during insert to DB", err)
	}

	return err
}

func (m *Manager) GetDB() *mgo.Session{
	return m.session
}

func (m *Manager) Query(query string) (string, error) {
	// TODO query DB
	//col := m.session.DB(database).C("")
	//"{"timestamp":{$gte: "2019-11-13", $lt: "2019-11-14"}}"

	return "", nil
}