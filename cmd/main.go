package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Key struct {
	ID  json.Number `json:"id"`
	Key string      `json:"key"`
}

func getKeys(user string) ([]Key, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/keys", user)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var keyResult []Key
	err = json.NewDecoder(resp.Body).Decode(&keyResult)
	if err != nil {
		return nil, err
	}
	return keyResult, nil
}

func saveKeys(keys []Key, authorizedKeyFile string) error {
	file, err := os.Create(authorizedKeyFile)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, key := range keys {
		_, err := file.WriteString(fmt.Sprintf("%s\n", key.Key))
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	user, exists := os.LookupEnv("GITHUB_USER")
	if !exists {
		fmt.Println("missing environment variable GITHUB_USER")
		os.Exit(1)
	}

	authorizedKeyFile, exists := os.LookupEnv("AUTHORIZED_HOSTS")
	if !exists {
		fmt.Println("missing environment variable AUTHORIZED_HOSTS")
		os.Exit(2)
	}

	keys, err := getKeys(user)
	if err != nil {
		fmt.Printf("failed to get list of GitHub keys. Error: %s\n", err)
		os.Exit(3)
	}

	err = saveKeys(keys, authorizedKeyFile)
	if err != nil {
		fmt.Printf("failed to get save keys to file. Error: %s\n", err)
		os.Exit(4)
	}
}
