package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

const TODO = "TODO"
const DOING = "DOING"
const DONE = "DONE"

type Config struct {
  path string
}

type Meeting struct {
  text string
  hour string
  minutes string
  state string
}

type Notification struct {
  text string
}

// read this from a config file in the future
func GetConfig() Config {
  return Config{path: "/home/gustavo/notes/"}
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

func ParseLineLookingForMeeting(line string) (bool, Meeting) {
  re :=regexp.MustCompile(`^.*-\s\[(?P<status>\s|x|-)\]\s(?P<hour>\d\d):(?P<minute>\d\d)\s(?P<text>.+)$`)
  matches := re.FindStringSubmatch(line)
  if len(matches) == 0 {
    return false, Meeting{}
  }

  result := make(map[string]string)
  for i, name := range re.SubexpNames() {
    if i != 0 && name != "" {
      result[name] = matches[i]
    }
  }

  var state string
  switch result["status"] {
  case " ":
    state = TODO
  case "x":
    state = DONE
  case "-":
    state = DOING
  default:
    state = DONE
  }

  return true, Meeting{
    text: result["text"],
    hour: result["hour"],
    minutes: result["minute"],
    state: state,
  }
}

func ParseMeetings(journal string) []Meeting {
  lines := strings.Split(journal, "\n")
  var meetings []Meeting
  for _, line := range lines {
    found, meeting := ParseLineLookingForMeeting(line)
    if found {
      meetings = append(meetings, meeting)
    }
  }

  return meetings
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
  config := GetConfig()
  filename := AppendMarkdownExtension(GetTodayDateAsString())
  filepath := GetFilePath(config.path, filename)
  filedata, err := os.ReadFile(filepath)
  if err != nil {
    panic(err)
  }
  journal := string(filedata)
  meetings := ParseMeetings(journal)
  now := time.Now()
  hour := fmt.Sprintf("%02d", now.Local().Hour())
  minutes := fmt.Sprintf("%02d", now.Local().Minute())

  fmt.Printf("meetings today: %q", meetings)
  fmt.Printf("now: %s:%s\n", hour, minutes)

  for _, meeting := range meetings {
    if hour == meeting.hour && minutes == meeting.minutes {
      notification := Notification{ text: meeting.text }
      NotifyGnome(notification)
      NotifyTmux(notification)
    }
  }
}
