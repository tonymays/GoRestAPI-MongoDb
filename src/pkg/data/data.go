package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type Data struct {
	Config			Configuration
	MongoClient		*mongo.Client
}

type Configuration struct {
	DbVendor			string
	DbName				string
	MongoUri			string
	Secret				string
    HTTPS				string
    Cert				string
    Key					string
    ServerListenPort	string
	RootUserName		string
	RootPassword		string
	Address				string
	City				string
	State				string
	Zip					string
	Country				string
	Email				string
	Phone				string
}

func (rcvr *Data) Init(e string) error {
	var config Configuration
	config.Init(e)
	rcvr.Config = config
	if config.DbVendor == "MongoDb" {
		mongoClient, err := ConnectMongoDb(rcvr.Config.MongoUri)
		if err != nil {
			return err
		}
		rcvr.MongoClient = mongoClient
	}
	return nil
}

func (rcvr *Configuration) Init(e string) error {
	confFile := "conf.json"
	if e == "test" {
		confFile = "conf_test.json"
	}
	file, _ := os.Open(confFile)
	decoder	:= json.NewDecoder(file)
	settings := Configuration{}
	err	:= decoder.Decode(&settings)
	if err != nil {
		return err
	}
	return nil
}
