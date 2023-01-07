package core

import (
  "regexp"
  "strings"
)

const TODO = "TODO"
const DOING = "DOING"
const DONE = "DONE"

const lineMeetingRegex = `^.*-\s\[(?P<status>\s|x|-)\]\s(?P<hour>\d\d):(?P<minute>\d\d)\s(?P<text>.+)$`

type Meeting struct {
  text string
  hour string
  minutes string
  state string
}

func NewMeeting(text, hour, minutes, state string) *Meeting {
  return &Meeting{
    text: text,
    hour: hour,
    minutes: minutes,
    state: state,
  }
}

func GetText(meeting *Meeting) string {
  return meeting.text
}

func GetHour(meeting *Meeting) string {
  return meeting.hour
}

func GetMinutes(meeting *Meeting) string {
  return meeting.minutes
}

func GetState(meeting *Meeting) string {
  return meeting.state
}

func ParseLineLookingForMeeting(line string) (bool, *Meeting) {
  re := regexp.MustCompile(lineMeetingRegex)
  matches := re.FindStringSubmatch(line)
  if len(matches) == 0 {
    return false, &Meeting{}
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

  return true, NewMeeting(result["text"], result["hour"], result["minute"], state)
}

func ParseMeetings(journal string) []*Meeting {
  lines := strings.Split(journal, "\n")
  var meetings []*Meeting
  for _, line := range lines {
    found, meeting := ParseLineLookingForMeeting(line)
    if found {
      meetings = append(meetings, meeting)
    }
  }

  return meetings
}
