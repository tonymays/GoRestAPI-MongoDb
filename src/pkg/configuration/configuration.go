package configuration

import (
	"encoding/json"
	"os"
)

// ---- Configuration structure ----
type Configuration struct {
	HTTPS				string
	Cert       			string
	Key        			string
	ServerListenPort	string
}

// ----
func CaptureAppConfiguration(env string) (Configuration, error) {
	// use a variable here siince we could have a test_conf.json file later
	confFile := "conf_test.json"
	if env == "production" {
		confFile = "conf.json"
	}

	file, _ := os.Open(confFile)
	decoder	:= json.NewDecoder(file)
	settings := Configuration{}
	err	:= decoder.Decode(&settings)

	// stop the train right here if we catch an error
	if err != nil {
		return Configuration{}, err
	}

	// ... otherwise, return the settings
	return settings, nil
}
