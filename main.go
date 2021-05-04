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
	TOKEN       = ""
	IconBitrise = "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"
)

var (
	APP = map[string]string{
		"": "",
	}
	JST = time.FixedZone("Asia/Tokyo", 9*60*60)
)

func main() {
	AppBuildList := makeAppBuildList()
	showTopInfo(AppBuildList)
	showRunningList(AppBuildList)
	showFinishedList(AppBuildList)
	showLastUpdate()
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

func (job *Job) BuildLink() string {
	return "https://app.bitrise.io/build/" + job.Id
}

func (job *Job) BuildStartTime() string {
	layout := "01/02 15:04"
	return job.StartAt.In(JST).Format(layout)
}

type ResponseData struct {
	Jobs []Job `json:"data"`
}

func getBuilds(appId string) ResponseData {
	url := "https://api.bitrise.io/v0.1/apps/" + appId + "/builds"
	now := time.Now()
	oneDayAgo := strconv.FormatInt(now.Add(-24*time.Hour).Unix(), 10)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token "+TOKEN)
	values := req.URL.Query()
	values.Set("after", oneDayAgo)
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

func makeAppBuildList() []AppBuilds {
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
	return AppBuildList
}

func showTopInfo(builds []AppBuilds) {
	runningTotal := 0
	for _, appBuild := range builds {
		runningTotal += len(appBuild.RunningJobs)
	}
	fmt.Println(strconv.Itoa(runningTotal) + " | image=" + IconBitrise)
	fmt.Println("---")
}

func showRunningList(builds []AppBuilds) {
	for _, app := range builds {
		fmt.Println(app.Name + ": " + strconv.Itoa(len(app.RunningJobs)) + " jobs")
		for _, job := range app.RunningJobs {
			fmt.Println("[" + job.Workflow + "]" + job.Branch + " : " + job.StatusText + " | href=" + job.BuildLink())
		}
		fmt.Println("---")
	}
}

func showFinishedList(builds []AppBuilds) {
	for _, app := range builds {
		fmt.Println("Last day: " + app.Name + ": " + strconv.Itoa(len(app.FinishedJobs)) + " jobs")
		for _, job := range app.FinishedJobs {
			fmt.Println(job.StatusEmoji() + " " + job.BuildStartTime() + " [" + app.Name + "/" + job.Workflow + "]" + job.Branch + " | href=" + job.BuildLink())
		}
		fmt.Println("---")
	}
}

func showLastUpdate() {
	now := time.Now().In(JST)
	fmt.Println("Last Updated: " + now.Format("2006-01-02 15:04:05") + " | disabled=true")
}
