package main

import (
	"log"
	"wow-api-go/auth"
)

func main() {
	clientCredential := auth.GetClientCredentials()
	log.Printf("%+v", clientCredential)
}
