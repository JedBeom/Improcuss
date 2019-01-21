package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	DB DB `json:"db"`
}

type DB struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var (
	config Config
)

func loadConfig(filename string) (config Config, err error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return
	}

	return
}

func loadConfigGlobal() {
	var err error
	config, err = loadConfig("config.json")
	if err != nil {
		log.Println("Error while loading config err:", err)
		os.Exit(1)
	}
}
