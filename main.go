package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"notes/config"
	"notes/core"
)

type Notification struct {
	text string
}

func GetTodayDateAsString() string {
	layout := "2006-01-02"
	now := time.Now()
	return fmt.Sprint(now.Format(layout))
}

func AppendMarkdownExtension(date string) string {
	return date + ".md"
}

func GetFilePath(notesPath string, filename string) string {
	return notesPath + "journal/" + filename
}

func NotifyGnome(notification Notification) {
	err := exec.Command("notify-send", "-u", "critical", "Atenção: Reunião", fmt.Sprintf("%s", notification.text)).Run()
	if err != nil {
		panic(err)
	}
}

func NotifyTmux(notification Notification) {
	err := exec.Command("tmux", "display-message", "-d", "5000", fmt.Sprintf("%s", notification.text)).Run()
	if err != nil {
		log.Println("Error displaying using tmux. Tmux is running?")
	}
}

func main() {
	cfg := config.GetConfig()
	filename := AppendMarkdownExtension(GetTodayDateAsString())
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
			notification := Notification{text: text}
			NotifyGnome(notification)
			NotifyTmux(notification)
		}
	}
}
