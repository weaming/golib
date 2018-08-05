package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// not validated
func NewConfig(filepath string, config *interface{}) *interface{} {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalln("Read config file error.")
	}
	err = json.Unmarshal(file, config)
	if err != nil {
		panic(err)
	}

	return config
}
