package pipedrive

import (
	"encoding/json"
	"errors"
	"fmt"
	ctoai "github.com/cto-ai/sdk-go"
	"github.com/grokify/html-strip-tags-go" // => strip
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type ResponseBody struct {
	Success        bool           `json:"success"`
	Data           []Note     `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
	Error          string         `json:"error"`
	ErrorCode      int            `json:"errorCode"`
}

type Note struct {
	Content      string       `json:"content"`
	AddTime      string       `json:"add_time"`
	UpdateTime   string       `json:"update_time"`
	Organization Organization `json:"organization"`
	User         User         `json:"user"`
	UserEmail    string       `json:"user.email"`
}

type Organization struct {
	Name string `json:"name"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AdditionalData struct {
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Start                int  `json:"start"`
	Limit                int  `json:"limit"`
	MoreItemInCollection bool `json:"more_items_in_collection"`
}

func getPipedriveNotes(client *ctoai.Client, apikey string, response *ResponseBody) {
	defaultDate := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	resp, err := client.Prompt.Input("forDate", "What day would you like to run it on? (YYYY-MM-DD)", ctoai.OptInputDefault(defaultDate))
	if err != nil {
		panic(err)
	}

	//we want to validate the date
	forDate, err := time.Parse("2006-01-02", resp)
	if err != nil {
		panic(errors.New("invalid date entered"))
	}

	if err = callNotesAPI(apikey, forDate.Format("2006-01-02"), response); err != nil {
		panic(err)
	}

}

func callNotesAPI(apikey string, forDate string, body *ResponseBody) error {
	url := "https://api.pipedrive.com/v1/notes"

	//Setup the Request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	//Add the Query Params
	q := request.URL.Query()
	q.Add("api_token", apikey)
	q.Add("start_date", forDate)
	q.Add("end_date", forDate)
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
	if err = json.Unmarshal(buf, &body); err != nil {
		panic(err)
	}

	if true != body.Success {
		return errors.New(fmt.Sprintf("Error!  Pipedrive Response Code %d : %s", body.ErrorCode, body.Error))
	}

	return nil

}

func printNotes(response *ResponseBody, client *ctoai.Client) {

	if len(response.Data) > 0 {
		fromTz, err := time.LoadLocation("UTC")
		if err != nil {
			panic(err)
		}

		toTz, err := time.LoadLocation("EST")
		if err != nil {
			panic(err)
		}

		for _, item := range response.Data {
			timeOfCall := fmt.Sprintf("%s UTC", item.UpdateTime)
			timeWithUTC := fmt.Sprintf("%s MST", item.UpdateTime)
			time, err := time.ParseInLocation("2006-01-02 15:04:05 MST", timeWithUTC, fromTz)
			if err == nil {
				time = time.In(toTz)
				timeOfCall = time.Format("2006-01-02 15:04 MST")
			}

			line1 := fmt.Sprintf("%s ( %s ) at %s", item.User.Name, item.User.Email, timeOfCall)
			line2 := fmt.Sprintf("With: %s", item.Organization.Name)

			line3 := strings.Replace(item.Content, "<br />", "\n", -1)
			line3 = strings.Replace(line3, "<br>", "\n", -1)
			line3 = strings.Replace(line3, "&nbsp;", " ", -1)
			line3 = fmt.Sprintf("Notes: %s", strip.StripTags(line3))

			if err = client.Ux.Print(line1); err != nil {
				panic(err)
			}

			if err = client.Ux.Print(line2); err != nil {
				panic(err)
			}

			if err = client.Ux.Print(line3); err != nil {
				panic(err)
			}
		}
	}
}
