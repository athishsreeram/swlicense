package servicetodomain

import (
	"log"
	"output/es-cud-output/client/nats/pub"

	"github.com/Jeffail/gabs"
	cfg "output/es-cud-output/config"
	"output/es-cud-output/domain"
)

func JWTCommandToEvent(data string) {

	jsonObj, _ := gabs.ParseJSON([]byte(data))

	key, msg := jsonObj.Search("command").Data().(string), data
	log.Println("%s %s", key, msg)
	domain.CreateCommand(string(jsonObj.Bytes()))

	if key == "JWTTokenCreateCommand" {

		pub.SendMsgStr(cfg.Conf.NATSSubj, "JWTDomainsEventCreated", "event", data)

	}

	if key == "JWTTokenUpdateCommand" {

		pub.SendMsgStr(cfg.Conf.NATSSubj, "JWTDomainsEventUpdated", "event", data)

	}

}
