package main

import (
	"fmt"
	"time"
)

func main() {
	startAt, _ := time.Parse("2006-01-02", "2021-04-30")
	job := Job{"123", "iOS", "Push", "main", startAt, "success"}
	fmt.Println(job.Id)
	fmt.Println(job.StatusEmoji())
	fmt.Println(job.StartAt)
}

type Job struct {
	Id string
	Platform string
	Workflow string
	Branch string
	StartAt time.Time
	Status string
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