package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// TODO these should be protected
var clientID string
var clientSecret string

func main() {

	// Build authURL
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

			// Get Access Token
			if code != "" {
				responseToken := requestFetchAccessToken(code)

				fmt.Printf("responseToken => %s  \n", responseToken)
				onedriveItems := getOneDriveRootDir(responseToken.AccessToken)
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
		"client_id":     clientID,
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

// TODO Tentative models

///////////////

// Token
type ResponseToken struct {
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    string `json:"expires_in"`
	ExtExpiresIn string `json:"ext_expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

///////////////

type OneDriveIdentity struct {
	DisplayName string `json:"displayName"`
	ID          string `json:"id"`
}
type OneDriveIdentitySet struct {
	Application OneDriveIdentity `json:"application"`
	Device      OneDriveIdentity `json:"device"`
	Group       OneDriveIdentity `json:"group"`
	User        OneDriveIdentity `json:"user"`
}

type OneDriveItemReference struct {
	DriveID   string `json:"driveId"`
	DriveType string `json:"driveType"`
	ID        string `json:"id"`
	ListID    string `json:"listId"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	ShareID   string `json:"shareId"`
	SiteID    string `json:"siteId"`
}

type OneDriveBaseItem struct {
	ID                   string                `json:"id"`
	CreatedBy            OneDriveIdentitySet   `json:"createdBy"`
	CreatedDateTime      string                `json:"createdDateTime"`
	Description          string                `json:"description"`
	ETag                 string                `json:"eTag"`
	LastModifiedBy       OneDriveIdentitySet   `json:"lastModifiedBy"`
	LastModifiedDateTime string                `json:"lastModifiedDateTime"`
	Name                 string                `json:"name"`
	ParentReference      OneDriveItemReference `json:"parentReference"`
	WebURL               string                `json:"webUrl"`
}
type OneDriveDriveItems struct {
	Value []OneDriveDriveItem `json:"value"`
}
type OneDriveDriveItem struct {
	/* inherited from baseItem */
	DriveID   string `json:"driveId"`
	DriveType string `json:"driveType"`
	ID        string `json:"id"`
	ListID    string `json:"listId"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	ShareID   string `json:"shareId"`
	SiteID    string `json:"siteId"`
	/* */
	Folder OneDriveFolder `json:"folder"`
	/* */
	File OneDriveFile `json:"file"`
}
type OneDriveFolder struct {
	ChildCount int `json:"childCount"`
}
type OneDriveFile struct {
	Hashes   OneDriveHashes `json:"hashes"`
	MimeType string         `json:"mimeType"`
}

type OneDriveHashes struct {
	Crc32Hash    string `json:"crc32Hash"`
	Sha1Hash     string `json:"sha1Hash"`
	QuickXorHash string `json:"quickXorHash"`
}
