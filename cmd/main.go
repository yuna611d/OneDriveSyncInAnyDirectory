package main

import (
	"flag"
	"fmt"
	"os"

	// TODO Tentative
	"./auth"
	"./server"
)


// TODO Remind: I need use Azure app Applications from personal account

func main() {

	// Arguments
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 || args[0] == "" || args[1] == "" {
		fmt.Printf("client and client secret are needed... \n")
		fmt.Printf("Program is ended \n")
		os.Exit(1)
	}

	clientID := args[0]
	clientSecret := args[1]
	tenantID := "common"
	fmt.Printf("tenant_id is %s \n", tenantID)
	fmt.Printf("client_id is %s \n", clientID)
	fmt.Printf("client_secret is %s \n", clientSecret)
	// Instansiate auth information
	authInfo := auth.GetInstance()
	authInfo.TenantID = tenantID
	authInfo.ClientID = clientID
	authInfo.ClientSecret = clientSecret

	server.Run()

}
