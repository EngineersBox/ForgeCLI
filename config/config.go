package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func ReadConfig() ForgeCLIConfig {
	configDir := "forgeCliConfig.json"
	config := ForgeCLIConfig{}
	if _, err := os.Stat(configDir); err == nil {
		b, err := ioutil.ReadFile(configDir)
		check(err)
		err = json.Unmarshal(b, &config)
		check(err)
	} else {
		panic("No 'forgeCliConfig.json' configuration file found.")
	}
	return config
}
