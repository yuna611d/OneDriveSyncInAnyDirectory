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


func (self *authInfo) getBaseEndPoint() string {
	scheme := "https://"
	host := "login.microsoftonline.com"
	baseEndPoint :="/oauth2/v2.0"
	return scheme + host + "/" + self.TenantID + baseEndPoint
}
func (self *authInfo) GetAuthEndPoint() string {
	baseEndPoint := self.getBaseEndPoint()
	return baseEndPoint + "/authorize"
}
func (self *authInfo) GetTokenEndPoint() string {
	baseEndPoint := self.getBaseEndPoint()
	return baseEndPoint + "/token"	
}



// Auth Flow
// https://docs.microsoft.com/ja-JP/azure/active-directory/develop/v2-oauth2-auth-code-flow


func (self *authInfo) GetAuthCodeRequestURI() string {

	endPoint :=  self.GetAuthEndPoint()

	queryMap := map[string]string{
		"client_id":     self.ClientID,
		// "scope":         "files.readwrite offline_access",
		"scope":         "Files.ReadWrite.All offline_access",
		"response_type": "code",
		"redirect_uri":  "http://localhost:5001/",
	}

	isInitialParameter := true
	authURL := endPoint
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

	fmt.Printf("Requested AuthURL is %s \n", authURL)

	return authURL
}

func (self *authInfo) RequestAccessToken() {

	endPoint := self.GetTokenEndPoint()

	// Build Body parameters
	bodyParameters := map[string]string{
		"client_id":     self.ClientID,
		"client_secret": self.ClientSecret,
		"code":          self.Code,
		"scope":         "Files.ReadWrite.All offline_access",
		"grant_type":    "authorization_code",
		"redirect_uri":  "http://localhost:5001/",
	}
	val := url.Values{}
	for key, value := range bodyParameters {
		val.Add(key, value)
	}
	fmt.Printf("Requested Token EndPont is %s \n", endPoint)
	fmt.Printf("Post Body values %s \n", val)

	client := http.Client{}
	req, err := http.NewRequest("POST", endPoint, strings.NewReader(val.Encode()))
	if err != nil {
		// TODO error handle
		// fmt.Printf("err => %", err.Error())
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
	if err := json.NewDecoder(res.Body).Decode(&responseToken); err != nil {
		fmt.Printf("decode error of responseToken: %s  \n", err)
	}
	// getJSON(res.Body, &responseToken)

	self.AccessToken = responseToken.AccessToken
	self.RefreshToken = responseToken.RefreshToken
	fmt.Printf(" ResponseToken => %s  \n", responseToken)
	fmt.Printf("  TokenType     => %s  \n", responseToken.TokenType)
	fmt.Printf("  Scope         => %s  \n", responseToken.Scope)
	fmt.Printf("  AccessToken   => %s  \n", responseToken.AccessToken)
	fmt.Printf("  RefreshToken  => %s  \n", responseToken.RefreshToken)
}
