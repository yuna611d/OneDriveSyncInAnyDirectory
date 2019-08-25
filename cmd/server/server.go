package server

import (
	"fmt"
	"net/http"

	// TODO tentative
	"../api"
	"../auth"
)

func Run() {

	// Http Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// See auth flow https://docs.microsoft.com/ja-jp/onedrive/developer/media/authorization_code_flow.png?view=odsp-graph-online

		// TODO Tentative
		if r.URL.Query().Get("code") == "" {
			link := "<a href=\"" + auth.GetInstance().GetAuthCodeRequestURI() + "\">Auth Request</a>"
			fmt.Fprintf(w, link)
		} else {
			// Redirected from Authorization end point

			// TODO message should be controlled from outside
			fmt.Fprintln(w, "Authorized")

			// Auth Code
			queryMap := r.URL.Query()
			code := queryMap.Get("code")
			// TODO work in 2nd times?
			// TODO Test for check code
			fmt.Printf("Acquired auth code is %s \n", code)
			auth.GetInstance().Code = code

			// Get Access Token
			if code != "" {

				auth.GetInstance().RequestAccessToken()

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
