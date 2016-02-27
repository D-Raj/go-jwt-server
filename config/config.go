/*
Configuration package is used to read the configuration file
config.json which stores the server port for current implementation
    {
        "ServerPort": ":8081"
    }
*/
package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//Stores the main configuration for the application
type Configuration struct {
	DB_HOST string
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT int
	DB_NAME string
}

var err error
var config Configuration

//ReadConfig will read the configuration json file to read the parameters
//which will be passed in the config file
func ReadConfig(fileName string) Configuration {

	configFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Unable to read config file '%s'", fileName)
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Print(err)
	}

	return config
}
