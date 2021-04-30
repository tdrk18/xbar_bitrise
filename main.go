package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	TOKEN = ""
)

var (
	APP = map[string]string{
		"": "",
	}
)

func main() {
	for name, id := range APP {
		data := getBuilds(id)
		fmt.Println(name)
		for _, job := range data.Jobs {
			fmt.Println(job.Id)
			fmt.Println(job.StatusEmoji())
			fmt.Println(job.StartAt)
		}
	}
}

type Job struct {
	Id string `json:"slug"`
	Workflow string `json:"triggered_workflow"`
	Branch string `json:"branch"`
	StartAt time.Time `json:"triggered_at"`
	Status string `json:"status_text"`
}

func (job *Job) StatusEmoji() string {
	switch job.Status {
	case "success":
		return ":ok:"
	case "error":
		return ":rotating_light:"
	case "aborted":
		return ":hand:"
	default:
		return ""
	}
}

type ResponseData struct {
	Jobs []Job `json:"data"`
}

func getBuilds(appId string) ResponseData {
	url := "https://api.bitrise.io/v0.1/apps/" + appId + "/builds"
	now := time.Now()
	oneHourAgo := strconv.FormatInt(now.Add(-1 * time.Hour).Unix(), 10)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token " + TOKEN)
	values := req.URL.Query()
	values.Set("after", oneHourAgo)
	req.URL.RawQuery = values.Encode()

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	var data ResponseData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}