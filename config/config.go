package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

const defaultConfigPath = "/.config/config.json"

type config struct {
	Evaneos evaneosConfig `json:"evaneos"`
}
type evaneosConfig struct {
	Slack slackConfig `json:"slack"`
}
type slackConfig struct {
	Token string `json:"token"`
}

func getUser() string {
	user, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	return user.HomeDir
}

// GetToken todo
func GetToken() string {
	configPath := getUser() + defaultConfigPath
	log.Printf("Using config file \"%s\"", configPath)

	// Open our JSON file
	jsonFile, err := os.Open(configPath)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer jsonFile.Close()

	// Inject JSON into type structures
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config config
	json.Unmarshal(byteValue, &config)

	if "" == config.Evaneos.Slack.Token {
		log.Fatal(fmt.Sprintf("The following JSON structure is required for %s: .evaneos.slack.token", configPath))
	}

	return config.Evaneos.Slack.Token
}
