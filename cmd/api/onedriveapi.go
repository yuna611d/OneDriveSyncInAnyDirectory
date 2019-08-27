package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"io/ioutil"

	// TODO tentative
	"../models"
	"../auth"
)

// TODO is it possible to use generics?
func callOneDriveAPI(apiURI string) io.ReadCloser {
	client := http.Client{}
	req, err := http.NewRequest("GET", apiURI, nil)
	if err != nil {
		// TODO error handle
		fmt.Printf("request for api is failed \n")
		os.Exit(1)
	}
	req.Header.Set("Authorization", "bearer "+ auth.GetInstance().AccessToken)
	res, err := client.Do(req)
	if err != nil {
		// TODO error handle
		fmt.Printf("response from api is failed \n")
		os.Exit(1)
	}
	defer res.Body.Close()

	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		fmt.Println("err " )		
	}
	fmt.Println("[body] " + string(body))


	return res.Body
}

func GetOneDriveRootDir() models.OneDriveDriveItems {
	// apiURI := "https://graph.microsoft.com/v1.0/users/me/drive/root/childrens"
	apiURI := "https://graph.microsoft.com/v1.0/me/drive/root/children"

	client := http.Client{}
	req, err := http.NewRequest("GET", apiURI, nil)
	if err != nil {
		// TODO error handle
		fmt.Printf("request for api is failed \n")
		os.Exit(1)
	}
	req.Header.Set("Authorization", "bearer "+ auth.GetInstance().AccessToken)
	res, err := client.Do(req)
	if err != nil {
		// TODO error handle
		fmt.Printf("response from api is failed \n")
		os.Exit(1)
	}
	defer res.Body.Close()


	target := models.OneDriveDriveItems{}
	// TODO tentative
	// var d struct {
	// 	Value []struct {
	// 		Name string `json:"name"`
	// 		Size int    `json:"size,omitempty"`
	// 		File struct {
	// 		} `json:"file,omitempty"`
	// 		Folder struct {
	// 			ChildCount int `json:"childCount"`
	// 		} `json:"folder,omitempty"`
	// 	} `json:"value"`
	// 	// OdataNextLink string `json:"@odata.nextLink"`
	// }
	// getJSON(res.Body, &d)
	// fmt.Printf("onedriveItem => %s", d)

	getJSON(res.Body, &target)



	return target
}

func getJSON(reader io.Reader, target interface{}) {
	if err:= json.NewDecoder(reader).Decode(target); err != nil {
		fmt.Printf("decode error: %s  \n", err)
	}
}
