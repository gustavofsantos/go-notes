package main

import (
	"fmt"
	"os"
	"time"

	"notes/config"
	"notes/core"
)

func GetTodayDateAsString(now time.Time) string {
	layout := "2006-01-02"
	return fmt.Sprintf("%s.md", now.Format(layout))
}

func AppendMarkdownExtension(date string) string {
	return date + ".md"
}

func GetFilePath(notesPath string, filename string) string {
	return notesPath + "journal/" + filename
}


func main() {
	cfg := config.GetConfig()
	filename := AppendMarkdownExtension(GetTodayDateAsString(time.Now()))
	filepath := GetFilePath(config.GetPath(cfg), filename)
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
