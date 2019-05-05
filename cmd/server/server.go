package server

import (
	"fmt"
	"net/http"

	// TODO tentative
	"../api"
	"../models"
)

var _authInfo models.AuthInfo

func SetAuthInfo(authInfo models.AuthInfo) {
	_authInfo = authInfo
}

func Run() {
	authURL := getAuthRequestURI()
	fmt.Printf("authURL is %s \n", authURL)

	// Http Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// TODO Tentative
		if r.URL.Query().Get("code") == "" {
			link := "<a href=\"" + authURL + "\">Auth Request</a>"
			fmt.Fprintf(w, link)
		} else {
			// TODO message should be controller from outside
			fmt.Fprintln(w, "Authorized")

			// Auth Code
			queryMap := r.URL.Query()
			code := queryMap.Get("code")
			// TODO work in 2nd times?
			// TODO Test for check code
			fmt.Printf("Acquired code is %s \n", code)
			_authInfo.Code = code

			// Get Access Token
			if code != "" {

				api.SetAuthInfo(_authInfo)
				api.RequestFetchAccessToken()
				onedriveItems := api.GetOneDriveRootDir()
				fmt.Printf("onedriveItem => %s", onedriveItems)

			}
		}
	})

	// TODO port should be specified from outside
	port := ":5001"
	fmt.Printf("Server is listening at %s \n", port)
	// TODO This should be https
	http.ListenAndServe(port, nil)

}

func getAuthRequestURI() string {
	// TODO These query parameteres should be configured from outside
	baseURL := "https://login.microsoftonline.com/common/oauth2/v2.0/authorize"
	queryMap := map[string]string{
		"client_id":     _authInfo.ClientID,
		"scope":         "files.readwrite offline_access",
		"response_type": "code",
		"redirect_uri":  "http://localhost:5001/",
	}

	isInitialParameter := true
	authURL := baseURL
	for key, value := range queryMap {
		operand := ""
		if isInitialParameter {
			operand = "?"
			isInitialParameter = !isInitialParameter
		} else {
			operand = "&"
		}
		authURL += operand + key + "=" + value
	}

	return authURL
}
