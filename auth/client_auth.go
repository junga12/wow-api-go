package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type ClientAuth struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func readClientAuth() *ClientAuth {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	authPath := path.Join(pwd, "auth", "client_auth.json")
	data, err := os.Open(authPath)
	if err != nil {
		panic(err)
	}

	var auth ClientAuth
	byteValue, _ := ioutil.ReadAll(data)
	err = json.Unmarshal(byteValue, &auth)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &auth
}
