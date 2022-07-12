package auth

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const authServer = "https://kr.battle.net/oauth/token"

type ClientCredential struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   uint   `json:"expires_in"`
	Sub         string `json:"sub"`
}

func GetClientCredentials() ClientCredential {
	clientAuth := readClientAuth()

	const grantType = "client_credentials"
	grantTypeBody := "grant_type=" + grantType
	req, err := http.NewRequest(http.MethodPost, authServer, bytes.NewBuffer([]byte(grantTypeBody)))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(clientAuth.ClientId, clientAuth.ClientSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	respBodyBytes, _ := ioutil.ReadAll(resp.Body)
	var clientCredential ClientCredential
	err = json.Unmarshal(respBodyBytes, &clientCredential)
	if err != nil {
		log.Fatalln(err)
	}
	return clientCredential
}
