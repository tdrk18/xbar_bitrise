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
	IconBitrise = "iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAAXNSR0IArs4c6QAAAAlwSFlzAAAOJgAADiYBou8l/AAAActpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IlhNUCBDb3JlIDUuNC4wIj4KICAgPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4KICAgICAgPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIKICAgICAgICAgICAgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIgogICAgICAgICAgICB4bWxuczp0aWZmPSJodHRwOi8vbnMuYWRvYmUuY29tL3RpZmYvMS4wLyI+CiAgICAgICAgIDx4bXA6Q3JlYXRvclRvb2w+d3d3Lmlua3NjYXBlLm9yZzwveG1wOkNyZWF0b3JUb29sPgogICAgICAgICA8dGlmZjpPcmllbnRhdGlvbj4xPC90aWZmOk9yaWVudGF0aW9uPgogICAgICA8L3JkZjpEZXNjcmlwdGlvbj4KICAgPC9yZGY6UkRGPgo8L3g6eG1wbWV0YT4KGMtVWAAAAl5JREFUOBF1U7+PUkEQnn0/IAZ4cIEzBDHEXGNyF73q0JaGYGltYS6x0tKY6B8gsTZqYaQytDQmeDkhZyOJ1BYUZ6NcwUGAAIHHe7x1voU1p8ZJPubHzs5+M/Mw6V8xOCSz2ewdJxqpOk7sQcyJjyeTyTeOq7OLV6yNI1hfROC68x3TMBMcnwcr7/omz9xoKAmgIgROwFgxfAYdHNz6ktpOLVKpbfvG/n4TMRaPgTwAd/4QzQBFYUOeM14riwiv6xx9vk5MJhJ3M1evPEkktiYUrFlJIZYr39+RJC3LsjtSyrAhREBC0Hg8inR//Hw1GA7fqxkMRqP020olXygUyHVdzhFkGAb5vk98kWzbvhYEgbLD4TC1Wi0qlUofwEwPMZ5MJikej7sc0y2gRz00zAa01aw4N8z2JQYZN/f2Dlk/nk6n8O3VamXxyygM2wD4dZuZWJ7nIR6azWbIfXQ7nz80Ft7yPjvHTA2vGEdHH6lWq8n5fE6maep2ZKfToUqlIrvdrohGo0vO/TqZTe9Z3C5op9Ajv0a7u3vQAjYEM2AIx3GoWCxSJBJRc+KjLSGpb8Vi8RfsvONR84NmkMvldK9qVWAByWQy6B+xgHNDrNOX0+lnVrvd/sROuX9+/qbf78vFYoEkXoS6z+ZawIItiS30ej0EXzYajWO9BdFsNql7dqaogzZE1VibavwoyozE99NTHOuvGDY9rNfrfE8uWTyG/x94nON9PjlB2ae4qBlYoRDaIpv7g1bTV8bmR7OCq3Nh6wLDcrk8qFarY2yA439N4Pc/R2KovMoY5/RR4Beqzhqkr7rq2gAAAABJRU5ErkJggg=="
)

var (
	APP = map[string]string{
		"": "",
	}
	JST = time.FixedZone("Asia/Tokyo", 9*60*60)
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

	runningTotal := 0
	for _, appBuild := range AppBuildList {
		runningTotal += len(appBuild.RunningJobs)
	}
	fmt.Println(strconv.Itoa(runningTotal) + " | image=" + IconBitrise)
	fmt.Println("---")

	for _, app := range AppBuildList {
		fmt.Println(app.Name + ": " + strconv.Itoa(len(app.RunningJobs)) + " jobs")
		for _, job := range app.RunningJobs {
			fmt.Println("[" + job.Workflow + "]" + job.Branch + " : " + job.StatusText + " | href=" + job.BuildLink())
		}
		fmt.Println("---")
	}

	for _, app := range AppBuildList {
		fmt.Println("Last day: " + app.Name + ": " + strconv.Itoa(len(app.FinishedJobs)) + " jobs")
		for _, job := range app.FinishedJobs {
			fmt.Println(job.StatusEmoji() + " " + job.BuildStartTime() + " [" + app.Name + "/" + job.Workflow + "]" + job.Branch + " | href=" + job.BuildLink())
		}
		fmt.Println("---")
	}

	now := time.Now().In(JST)
	fmt.Println("Last Updated: " + now.Format("2006-01-02 15:04:05") + " | disabled=true")
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
