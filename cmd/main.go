package main

import (
	"flag"
	"fmt"
	"os"

	// TODO Tentative
	"./auth"
	"./server"
)

// TODO these should be protected
var clientID string
var clientSecret string

func main() {

	// Arguments
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 || args[0] == "" || args[1] == "" {
		fmt.Printf("Client Code is needed... \n")
		fmt.Printf("Usage client_code and client Secret is needed \n")
		fmt.Printf("Program is ended \n")
		os.Exit(1)
	}

	clientID = args[0]
	clientSecret = args[1]
	fmt.Printf("client_id is %s \n", clientID)
	fmt.Printf("client_secret is %s \n", clientSecret)
	// Instansiate auth information
	authInfo := auth.GetInstance()
	authInfo.ClientID = clientID
	authInfo.ClientSecret = clientSecret

	// Run Server to get access code
	// server.SetAuthInfo()
	server.Run()

}
