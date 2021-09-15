package configuration

import (
	"encoding/json"
	"os"
)

// ---- Configuration ----
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
	Firstname			string
	Lastname			string
	Address				string
	City				string
	State				string
	Zip					string
	Country				string
	Email				string
	Phone				string
}

// ---- Init ----
func Init(e string) (Configuration, error) {
	// prod or test
	confFile := "conf.json"
	if e == "test" {
		confFile = "conf_test.json"
	}

	// open the file and decode contents
	file, _ := os.Open(confFile)
	decoder	:= json.NewDecoder(file)

	// establish a nil struct and then populate it
	settings := Configuration{}
	err	:= decoder.Decode(&settings)
	if err != nil {
		return Configuration{}, err
	}

	// return the configuration settings at this point
	return settings, nil
}