package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type EVConfig struct {
	NetworkID uint64 `json:"network_id"`
	EVs []EV `json:"evs"`
}

type EV struct {
	Address string `json:"address"`
	Public  string `json:"public"`
	Private string `json:"private"`
}

func parseConfiguration(target string) EVConfig {
	jsonFile, err := os.Open(target)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully configuration file")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config EVConfig
	json.Unmarshal(byteValue, &config)

	return config
}
