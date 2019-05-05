package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func requestFetchAccessToken(code string) ResponseToken {
	tokenURL := "https://login.microsoftonline.com/common/oauth2/v2.0/token"

	// Build Body parameters
	bodyParameters := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          code,
		"grant_type":    "authorization_code",
		"redirect_uri":  "http://localhost:5001/",
	}
	val := url.Values{}
	for key, value := range bodyParameters {
		val.Add(key, value)
	}
	fmt.Printf("Post Body values %s \n", val)

	client := http.Client{}
	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(val.Encode()))
	if err != nil {
		// TODO error handle
		// fmt.Printf("err => %w", err.Error)
		fmt.Printf("request for access token is failed \n")
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		// TODO error handle
		fmt.Printf("response for access token is failed \n")
		os.Exit(1)
	}
	defer res.Body.Close()

	responseToken := ResponseToken{}
	getJSON(res.Body, &responseToken)
	return responseToken
}

// TODO is it possible to use generics?
func callOneDriveAPI(accessToken string, apiURI string) io.ReadCloser {
	client := http.Client{}
	req, err := http.NewRequest("GET", apiURI, nil)
	if err != nil {
		// TODO error handle
		fmt.Printf("request for api is failed \n")
		os.Exit(1)
	}
	req.Header.Set("Authorization", "bearer "+accessToken)
	res, err := client.Do(req)
	if err != nil {
		// TODO error handle
		fmt.Printf("response from api is failed \n")
		os.Exit(1)
	}
	defer res.Body.Close()
	return res.Body
}

func getOneDriveRootDir(accessToken string) OneDriveDriveItems {
	apiURI := "https://graph.microsoft.com/v1.0/users/me/drive/root/childrens"

	responseBody := callOneDriveAPI(accessToken, apiURI)
	// responseBody.Close()
	target := OneDriveDriveItems{}
	getJSON(responseBody, &target)
	return target
}

func getJSON(reader io.Reader, target interface{}) error {
	return json.NewDecoder(reader).Decode(target)
}
