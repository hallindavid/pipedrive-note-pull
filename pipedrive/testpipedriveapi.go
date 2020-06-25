package pipedrive

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type PipeDriveTestKeyResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	ErrorCode int    `json:"errorCode"`
}

func testApiKey(apiKey string)  error {
	url := "https://api.pipedrive.com/v1/users/me"

	//Setup the Request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	//Add the API KEY to Query Params
	q := request.URL.Query()
	q.Add("api_token", apiKey)
	request.URL.RawQuery = q.Encode()

	//Set some boundaries on the http client - we dont' want a long running process or anything here...
	client := &http.Client{
		Timeout: time.Second * 5,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	//Execute the request
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()


	//buf is now our response body as []byte
	buf, _ := ioutil.ReadAll(response.Body)

	//Parse the response as a JSON object (what we'll actually use in prod)
	responseObj := PipeDriveTestKeyResponse{}
	if err = json.Unmarshal(buf, &responseObj); err != nil {
		panic(err)
	}

	if true != responseObj.Success {
		return errors.New(fmt.Sprintf("Error!  Pipedrive Response Code %d : %s", responseObj.ErrorCode, responseObj.Error))
	}

	return nil
}

