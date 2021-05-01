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
	var AppBuildList []AppBuilds
	for name, id := range APP {
		data := getBuilds(id)
		var runningList []Job
		var finishedList []Job
		for _, job := range data.Jobs {
			if job.Status == 0 {
				runningList = append(runningList, job)
			} else {
				finishedList = append(finishedList, job)
			}
		}
		AppBuildList = append(AppBuildList, AppBuilds{name, runningList, finishedList})
	}

	for _, app := range AppBuildList {
		fmt.Println("-- running --")
		fmt.Println(app.Name)
		for _, job := range app.RunningJobs {
			fmt.Println(job.Id)
			fmt.Println(job.Workflow)
			fmt.Println(job.StartAt.String())
			fmt.Println(job.StatusEmoji())
		}
		fmt.Println("-- finished --")
		fmt.Println(app.Name)
		for _, job := range app.FinishedJobs {
			fmt.Println(job.Id)
			fmt.Println(job.Workflow)
			fmt.Println(job.StartAt.String())
			fmt.Println(job.StatusEmoji())
		}
	}
}

type AppBuilds struct {
	Name         string
	RunningJobs  []Job
	FinishedJobs []Job
}

type Job struct {
	Id         string    `json:"slug"`
	Workflow   string    `json:"triggered_workflow"`
	Branch     string    `json:"branch"`
	StartAt    time.Time `json:"triggered_at"`
	Status     int       `json:"status"`
	StatusText string    `json:"status_text"`
}

func (job *Job) StatusEmoji() string {
	switch job.StatusText {
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
	oneHourAgo := strconv.FormatInt(now.Add(-1*time.Hour).Unix(), 10)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token "+TOKEN)
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
