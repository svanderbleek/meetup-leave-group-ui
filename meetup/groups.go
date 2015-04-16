package meetup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Group struct {
	Id   int
	Name string
}

type ApiListResponse struct {
	Results []Group
}

var apiKey = os.Getenv("MEETUP_API_KEY")
var apiBase = "https://api.meetup.com/2"
var apiLeave = apiBase + "/profile/%v/self?key=" + apiKey
var apiList = apiBase + "/groups?member_id=self&key=" + apiKey

func leaveUrl(id string) string {
	return fmt.Sprintf(apiLeave, id)
}

func LeaveGroup(id string) bool {
	url := leaveUrl(id)
	request, _ := http.NewRequest("DELETE", url, nil)
	response, _ := http.DefaultClient.Do(request)
	if response.StatusCode == 200 {
		return true
	} else {
		return false
	}
}

func listUrl() string {
	return apiList
}

func Groups() []Group {
	url := listUrl()
	response, _ := http.Get(url)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var list ApiListResponse
	json.Unmarshal(body, &list)
	return list.Results
}
