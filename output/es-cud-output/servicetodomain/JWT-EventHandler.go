package servicetodomain

import (
	"encoding/json"
	"log"
	"output/es-cud-output/client/nats/pub"
	cfg "output/es-cud-output/config"
	"output/es-cud-output/custom"
	"output/es-cud-output/domain"
	"output/es-cud-output/proto"

	"github.com/Jeffail/gabs"
)

func JWTEventProcesing(data string) {

	jsonObj, _ := gabs.ParseJSON([]byte(data))

	key, msg := jsonObj.Search("event").Data().(string), data
	log.Println("%s %s", key, msg)

	if key == "JWTTokenDomainsEventCreated" {
		pub.Send(cfg.Conf.NATSSubj, "JWTTokenDomainsEventCreateCompleted", "event", CreateProcessing(key, msg), jsonObj.Search("uuid").Data().(string))
	}

	if key == "JWTTokenDomainsEventUpdated" {
		pub.Send(cfg.Conf.NATSSubj, "JWTTokenDomainsEventUpdateCompleted", "event", UpdateProcessing(key, msg), jsonObj.Search("uuid").Data().(string))
	}

	if key == "JWTTokenDomainsEventReadd" {
		pub.Send(cfg.Conf.NATSSubj, "JWTTokenDomainsEventReadCompleted", "event", ReadProcessing(key, msg), jsonObj.Search("uuid").Data().(string))
	}

}

func CreateProcessing(key string, msg string) *proto.JWTToken {
	// TO-DO
	var user *proto.JWTRequest
	err := json.Unmarshal([]byte(msg), &user)
	if err != nil {
		log.Println(err)
	}

	log.Println("DTO Data", user.User)

	dom := custom.CreateToken(user.User)

	jWTTokenDomains := domain.ConvertJWTToken2JWTTokenDomains(dom)
	log.Println("Domain Data ", jWTTokenDomains)
	persistedDomain := domain.CreateJWTTokenDomains(jWTTokenDomains)
	jsonObj, _ := gabs.ParseJSON([]byte(msg))
	createEvent(jsonObj.Search("event").Data().(string), persistedDomain, jsonObj.Search("uuid").Data().(string), jsonObj.Search("command").Data().(string))
	JWTToken := domain.ConvertJWTTokenDomains2JWTToken(persistedDomain)
	return JWTToken

}

func UpdateProcessing(key string, msg string) *proto.JWTToken {
	// TO-DO
	var dat proto.JWTRenewRequest
	err := json.Unmarshal([]byte(msg), &dat)
	if err != nil {
		log.Println(err)
	}

	jWTTokenResp = domain.ConvertJWTTokenDomains2JWTToken(domain.ReadJWTTokenDomains(dat.JwtToken))

	dom := custom.CreateToken(jWTTokenResp.User)

	persistedDomain := domain.UpdateJWTTokenDomains(dom.User, dom)
	jsonObj, _ := gabs.ParseJSON([]byte(msg))
	createEvent(jsonObj.Search("event").Data().(string), persistedDomain, jsonObj.Search("uuid").Data().(string), jsonObj.Search("command").Data().(string))
	JWTToken := domain.ConvertJWTTokenDomains2JWTToken(persistedDomain)
	return JWTToken

}

func createEvent(event string, data interface{}, uuid string, cmd string) {
	msg, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	jsonObj, err := gabs.ParseJSON(msg)
	jsonObj.SetP(event, "event")
	jsonObj.SetP(uuid, "uuid")
	jsonObj.SetP(cmd, "command")
	log.Println(jsonObj)
	domain.CreateEvent(string(jsonObj.Bytes()))

}
