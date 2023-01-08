package core

import (
	"regexp"
	"strings"
)

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
