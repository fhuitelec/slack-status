package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

const defaultConfigPath = "/.slack-status.json"

type slackToken struct {
	Token string `json:"token"`
}

func getUser() string {
	user, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	return user.HomeDir
}

func askAndPersistToken(configPath string) {
	fmt.Print("Enter legacy slack token: ")
	var rawToken string
	fmt.Scanln(&rawToken)

	token := slackToken{rawToken}
	tokenJSON, _ := json.Marshal(token)

	ioutil.WriteFile(configPath, append(tokenJSON, []byte("\n")...), 0644)
}

// GetToken todo
func GetToken() string {
	configPath := getUser() + defaultConfigPath
	log.Printf("Using config file \"%s\"", configPath)

	// Open our JSON file
	jsonFile, err := os.Open(configPath)

	if err != nil && os.IsNotExist(err) {
		askAndPersistToken(configPath)
		defer jsonFile.Close()
		jsonFile, err = os.Open(configPath)
	} else if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer jsonFile.Close()

	// Inject JSON into type structures
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var slackToken slackToken
	json.Unmarshal(byteValue, &slackToken)

	if "" == slackToken.Token {
		log.Fatal(fmt.Sprintf("The following JSON structure is required for %s: .token", configPath))
	}

	return slackToken.Token
}
