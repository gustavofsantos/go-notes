package core

const TODO = "TODO"
const DOING = "DOING"
const DONE = "DONE"

const lineMeetingRegex = `^.*-\s\[(?P<status>\s|x|-)\]\s(?P<hour>\d\d):(?P<minute>\d\d)\s(?P<text>.+)$`

type Meeting struct {
	text    string
	hour    string
	minutes string
	state   string
}

func NewMeeting(text, hour, minutes, state string) *Meeting {
	return &Meeting{
		text:    text,
		hour:    hour,
		minutes: minutes,
		state:   state,
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
