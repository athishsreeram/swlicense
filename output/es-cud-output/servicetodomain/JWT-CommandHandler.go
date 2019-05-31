package servicetodomain

import (
	"log"
	"swlicense/output/es-cud-output/client/nats/pub"

	cfg "swlicense/output/es-cud-output/config"
	"swlicense/output/es-cud-output/domain"

	"github.com/Jeffail/gabs"
)

func JWTCommandToEvent(data string) {

	jsonObj, _ := gabs.ParseJSON([]byte(data))

	key, msg := jsonObj.Search("command").Data().(string), data
	log.Println("%s %s", key, msg)
	domain.CreateCommand(string(jsonObj.Bytes()))

	if key == "JWTTokenCreatedCommand" {

		pub.SendMsgStr(cfg.Conf.NATSSubj, "JWTTokenDomainsEventCreated", "event", data)

	}

	if key == "JWTTokenUpdatedCommand" {

		pub.SendMsgStr(cfg.Conf.NATSSubj, "JWTTokenDomainsEventUpdated", "event", data)

	}

}
