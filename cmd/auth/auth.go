package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"../models"
)

type authInfo struct {
	TenantID     string
	ClientID     string
	ClientSecret string
	Code         string
	AccessToken  string
	RefreshToken string
}

var instance *authInfo

func GetInstance() *authInfo {
	if instance == nil {
		instance = &authInfo{}
	}
	return instance
}

func (self *authInfo)authFlowURL() map[string] string {
	scheme := "https://"
	host := "login.microsoftonline.com"
	authorizePath := "/oauth2/v2.0/authorize"
	tokenPath := "/oauth2/v2.0/token"

	AuthURL := scheme + host + "/" + self.TenantID + authorizePath
	TokenURL := scheme + host + "/" + self.TenantID + tokenPath

	return map[string] string{
		"auth": AuthURL,
		"token": TokenURL,
	} 
}



// Auth Flow
// https://docs.microsoft.com/ja-JP/azure/active-directory/develop/v2-oauth2-auth-code-flow


func (self *authInfo) GetAuthCodeRequestURI() string {
	// TODO These query parameteres should be configured from outside
	scheme:="https://"
	host:="login.microsoftonline.com"
	path:=	"/oauth2/v2.0/authorize"

	baseURL :=  scheme + host + "/" +self.TenantID + path 
	queryMap := map[string]string{
		"client_id":     self.ClientID,
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

func (self *authInfo) RequestAccessToken() {
	scheme:="https://"
	host:="login.microsoftonline.com"
	path:=	"/oauth2/v2.0/token"

	tokenURL := scheme + host +  "/" +self.TenantID + path

	// Build Body parameters
	bodyParameters := map[string]string{
		"client_id":     self.ClientID,
		"client_secret": self.ClientSecret,
		"code":          self.Code,
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

	self.AccessToken = responseToken.AccessToken
	self.RefreshToken = responseToken.RefreshToken
	fmt.Printf("responseToken => %s  \n", responseToken)
	fmt.Printf("TokenType => %s  \n", responseToken.TokenType)
	fmt.Printf("Scope => %s  \n", responseToken.Scope)
	fmt.Printf("accessToken => %s  \n", responseToken.AccessToken)
	fmt.Printf("RefreshToken => %s  \n", responseToken.RefreshToken)

	// return responseToken
}
