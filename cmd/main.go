package main

import (
	"flag"
	"fmt"
	"os"

	// TODO Tentative
	"./auth"
	"./server"
)


func main() {

	// Arguments
	flag.Parse()
	args := flag.Args()
	if len(args) != 3 || args[0] == "" || args[1] == "" || args[2] == "" {
		fmt.Printf("tenant, client and client secret are needed... \n")
		fmt.Printf("Program is ended \n")
		os.Exit(1)
	}

	tenantID := args[0]
	clientID := args[1]
	clientSecret := args[2]
	fmt.Printf("tenant_id is %s \n", tenantID)
	fmt.Printf("client_id is %s \n", clientID)
	fmt.Printf("client_secret is %s \n", clientSecret)
	// Instansiate auth information
	authInfo := auth.GetInstance()
	authInfo.TenantID = tenantID
	authInfo.ClientID = clientID
	authInfo.ClientSecret = clientSecret

	// Run Server to get access code
	// server.SetAuthInfo()
	server.Run()

}
