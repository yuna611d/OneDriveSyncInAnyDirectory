package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	// TODO tentative
	"../models"
)

var _authInfo models.AuthInfo

func SetAuthInfo(authInfo models.AuthInfo) {
	_authInfo = authInfo
}

func RequestFetchAccessToken() {
	tokenURL := "https://login.microsoftonline.com/common/oauth2/v2.0/token"

	// Build Body parameters
	bodyParameters := map[string]string{
		"client_id":     _authInfo.ClientID,
		"client_secret": _authInfo.ClientSecret,
		"code":          _authInfo.Code,
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

	responseToken := models.ResponseToken{}
	json.NewDecoder(res.Body).Decode(&responseToken)
	// getJSON(res.Body, &responseToken)

	_authInfo.AccessToken = responseToken.AccessToken
	_authInfo.RefreshToken = responseToken.RefreshToken
	fmt.Printf("responseToken => %s  \n", responseToken)
	fmt.Printf("TokenType => %s  \n", responseToken.TokenType)
	fmt.Printf("Scope => %s  \n", responseToken.Scope)
	fmt.Printf("accessToken => %s  \n", responseToken.AccessToken)
	fmt.Printf("RefreshToken => %s  \n", responseToken.RefreshToken)

	// return responseToken
}

// TODO is it possible to use generics?
func callOneDriveAPI(apiURI string) io.ReadCloser {
	client := http.Client{}
	req, err := http.NewRequest("GET", apiURI, nil)
	if err != nil {
		// TODO error handle
		fmt.Printf("request for api is failed \n")
		os.Exit(1)
	}
	req.Header.Set("Authorization", "bearer "+_authInfo.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		// TODO error handle
		fmt.Printf("response from api is failed \n")
		os.Exit(1)
	}
	defer res.Body.Close()
	return res.Body
}

func GetOneDriveRootDir() models.OneDriveDriveItems {
	apiURI := "https://graph.microsoft.com/v1.0/users/me/drive/root/childrens"

	responseBody := callOneDriveAPI(apiURI)
	// responseBody.Close()
	target := models.OneDriveDriveItems{}
	getJSON(responseBody, &target)
	return target
}

func getJSON(reader io.Reader, target interface{}) {
	json.NewDecoder(reader).Decode(target)
}
