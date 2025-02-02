package domain

import (
	"log"

	"github.com/Jeffail/gabs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	cfg "swlicense/output/es-cud-output/config"
)

type CommandStores struct {
	Uuid    string
	Data    string
	Command string
}

func CreateCommand(data string) {

	jsonObj, _ := gabs.ParseJSON([]byte(data))
	var command string
	if jsonObj.ExistsP("command") == true {
		command = jsonObj.Search("command").Data().(string)
	}

	es := &CommandStores{
		Uuid:    jsonObj.Search("uuid").Data().(string),
		Data:    data,
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
