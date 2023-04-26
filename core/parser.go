package core

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

const lineMeetingRegex = `^.*-\s\[(?P<status>\s|x|-)\]\s(?P<hour>\d\d):(?P<minute>\d\d)\s(?P<text>.+)$`
const lineMeetingAltRegex = `^.*-\s(?P<status>TODO|DOING|DONE)\s(?P<hour>\d\d):(?P<minute>\d\d)\s(?P<text>.+)$` 

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

func ParseLineLookingForMeeting(line string) (bool, *Meeting) {
	re := regexp.MustCompile(lineMeetingRegex)
  reAlt := regexp.MustCompile(lineMeetingAltRegex)
	matches := re.FindStringSubmatch(line)
  matchesAlt := reAlt.FindStringSubmatch(line)
	if len(matches) == 0 && len(matchesAlt) == 0 {
		return false, &Meeting{}
	}

	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = matches[i]
		}
	}

  for i, name := range reAlt.SubexpNames() {
    if i != 0 && name != "" {
      result[name] = matchesAlt[i]
    }
  }

	state, err := parseMeetingState(result["state"])
	if err != nil {
		log.Printf(err.Error())
		return false, &Meeting{}
	}

	return true, NewMeeting(result["text"], result["hour"], result["minute"], state)
}

func parseMeetingState(state string) (MeetingState, error) {
	switch state {
	case " ":
		return TODO, nil
  case "TODO":
    return TODO, nil
	case "x":
		return DONE, nil
  case "DONE":
    return DONE, nil
	case "-":
		return DOING, nil
  case "DOING":
    return DOING, nil
	default:
		return "", fmt.Errorf("state %s is not recognized", state)
	}
}
