package main

import (
	"log"
	"wow-api-go/auth"
)

func main() {
	clientAuth := auth.GetClientAuth()
	log.Println(clientAuth.ClientId)
	log.Println(clientAuth.ClientSecret)
}
