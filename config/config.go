package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Name        string           `json:"name"`
	Version     string           `json:"version"`
	Environment string           `json:"environment"`
	Properties  []propertySource `json:"properties"`
}

type propertySource struct {
	Source map[string]interface{} `json:"source"`
}

func Init() {
	fetchConfigurations()
}

func fetchConfigurations() {
	bodyBytes, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		fmt.Println("Couldn't read configuration file.", err)
		return
	}

	var configs AppConfig
	err = json.Unmarshal(bodyBytes, &configs)
	if err != nil {
		fmt.Println("Cannot parse configuration, message: ", err.Error())
		return
	}

	for key, value := range configs.Properties[0].Source {
		viper.Set(key, value)
		fmt.Printf("Loading config property -> %s - %s \n", key, value)
	}
	fmt.Println("Successfully loaded configurations.")
}
