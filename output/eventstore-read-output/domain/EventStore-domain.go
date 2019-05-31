package domain

import (
	cfg "swlicense/output/eventstore-read-output/config"
	proto "swlicense/output/eventstore-read-output/proto"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	mapstructure "github.com/mitchellh/mapstructure"
	"log"
)

var engine *xorm.Engine

type EventStores struct {
	Uuid    string `mapstructure:"uuid"`
	Event   string `mapstructure:"event"`
	Command string `mapstructure:"command"`
	Data    string `mapstructure:"data"`
}

func ReadAllEventStores() []EventStores {
	var err error
	engine, err = xorm.NewEngine("mysql", cfg.Conf.DBCon)

	var eventStores []EventStores
	engine.Find(&eventStores)
	log.Println("{}", eventStores)

	if err != nil {
		log.Println(err)
	}

	return eventStores

}

func ReadEventStores(Uuid string) EventStores {
	var err error
	engine, err = xorm.NewEngine("mysql", cfg.Conf.DBCon)

	var eventStores = EventStores{Uuid: Uuid}
	has, err := engine.Get(&eventStores)
	log.Println("{}", eventStores)
	if err != nil {
		log.Println(err)
	}
	log.Println(has)

	return eventStores

}

func ConvertEventStores2EventStore(eventstores interface{}) (eventstore *proto.EventStore) {
	err0 := mapstructure.Decode(eventstores, &eventstore)
	if err0 != nil {
		panic(err0)
	}
	return eventstore
}

func ConvertEventStore2EventStores(eventstore *proto.EventStore) (eventstores EventStores) {
	err1 := mapstructure.Decode(eventstore, &eventstores)
	if err1 != nil {
		panic(err1)
	}
	return eventstores
}
