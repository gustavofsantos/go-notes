package main

import (
	"fmt"
	"os"
	"time"

	"notes/config"
	"notes/core"
)

func main() {
	cfg := config.GetConfig()
	filename := core.GetTodayDateAsString(time.Now())
	filepath := core.GetFilePath(config.GetPath(cfg), filename)
	filedata, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	journal := string(filedata)
	meetings := core.ParseMeetings(journal)
	now := time.Now()
	hour := fmt.Sprintf("%02d", now.Local().Hour())
	minutes := fmt.Sprintf("%02d", now.Local().Minute())

	for _, meeting := range meetings {
		meetingHour := core.GetHour(meeting)
		meetingMinutes := core.GetMinutes(meeting)
		text := core.GetText(meeting)
		if hour == meetingHour && minutes == meetingMinutes {
			notification := &Notification{text: text}
			HandleNotification(notification)
		}
	}
}
