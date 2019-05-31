package domain

import (
	"log"
	cfg "swlicense/output/es-cud-output/config"
	proto "swlicense/output/es-cud-output/proto"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	mapstructure "github.com/mitchellh/mapstructure"
)

var engine *xorm.Engine
var conn = "root:@tcp(localhost:3306)/GOGENETIC_SCHEMA?charset=utf8&parseTime=True&loc=Local"

type JWTTokenDomains struct {
	User  string `mapstructure:"User"`
	Token string `mapstructure:"Token"`
}

func ReadAllJWTTokenDomains() []JWTTokenDomains {
	var err error
	engine, err = xorm.NewEngine("mysql", cfg.Conf.DBCon)

	var jWTTokenDomains []JWTTokenDomains
	engine.Find(&jWTTokenDomains)
	log.Println("{}", jWTTokenDomains)

	if err != nil {
		log.Println(err)
	}

	return jWTTokenDomains

}

func ReadJWTTokenDomains(jwttoken string) JWTTokenDomains {
	var err error
	engine, err = xorm.NewEngine("mysql", cfg.Conf.DBCon)

	var jWTTokenDomains = JWTTokenDomains{Token: jwttoken}
	has, err := engine.Get(&jWTTokenDomains)
	log.Println("{}", jWTTokenDomains)
	if err != nil {
		log.Println(err)
	}
	log.Println(has)

	return jWTTokenDomains

}

func CreateJWTTokenDomains(jWTTokenDomains JWTTokenDomains) JWTTokenDomains {

	var err error
	engine, err = xorm.NewEngine("mysql", cfg.Conf.DBCon)
	engine.Insert(&jWTTokenDomains)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully Created {}", &jWTTokenDomains)
	}
	return jWTTokenDomains

}

func DeleteJWTTokenDomains(user string) JWTTokenDomains {
	var err error
	engine, err = xorm.NewEngine("mysql", cfg.Conf.DBCon)

	var jWTTokenDomains = JWTTokenDomains{User: user}
	engine.Delete(&jWTTokenDomains)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully Deleted {}", &jWTTokenDomains)
	}
	return jWTTokenDomains
}

func UpdateJWTTokenDomains(jWTTokenDomains JWTTokenDomains) JWTTokenDomains {
	var err error
	engine, err = xorm.NewEngine("mysql", cfg.Conf.DBCon)

	engine.Update(&jWTTokenDomains)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully Updated {}", &jWTTokenDomains)
	}

	return jWTTokenDomains
}

func ConvertJWTTokenDomains2JWTToken(jwttokendomains interface{}) (jwttoken *proto.JWTToken) {
	err0 := mapstructure.Decode(jwttokendomains, &jwttoken)
	if err0 != nil {
		panic(err0)
	}
	return jwttoken
}

func ConvertJWTToken2JWTTokenDomains(jwttoken *proto.JWTToken) (jwttokendomains JWTTokenDomains) {
	err1 := mapstructure.Decode(jwttoken, &jwttokendomains)
	if err1 != nil {
		panic(err1)
	}
	return jwttokendomains
}
