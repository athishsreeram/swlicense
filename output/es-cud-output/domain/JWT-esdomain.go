package domain

import (
	"log"

	"github.com/Jeffail/gabs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	cfg "output/es-cud-output/config"
)

type EventStores struct {
	Uuid    string
	Data    string
	Event   string
	Command string
}

func CreateEvent(data string) {

	jsonObj, _ := gabs.ParseJSON([]byte(data))
	var command string
	var event string
	if jsonObj.ExistsP("command") == true {
		command = jsonObj.Search("command").Data().(string)
	}

	if jsonObj.ExistsP("event") == true {
		event = jsonObj.Search("event").Data().(string)
	}

	es := &EventStores{
		Uuid:    jsonObj.Search("uuid").Data().(string),
		Data:    data,
		Event:   event,
		Command: command,
	}

	var err error
	engineEvent, err := xorm.NewEngine("mysql", cfg.Conf.DBCon)
	engineEvent.Insert(es)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully Created {}", es)
	}

}
