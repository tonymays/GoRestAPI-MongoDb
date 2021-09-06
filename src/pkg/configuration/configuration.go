package configuration

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	MongoUri			string
	DbName				string
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

func Init(e string) (Configuration, error) {
	confFile := "conf.json"
	if e == "test" {
		confFile = "conf_test.json"
	}

	file, _ := os.Open(confFile)
	decoder	:= json.NewDecoder(file)
	settings := Configuration{}
	err	:= decoder.Decode(&settings)
	if err != nil {
		return Configuration{}, err
	}
	return settings, nil
}