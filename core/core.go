package core

const TODO = "TODO"
const DOING = "DOING"
const DONE = "DONE"

type MeetingState string

type Meeting struct {
	text    string
	hour    string
	minutes string
	state   MeetingState
}

func NewMeeting(text, hour, minutes string, state MeetingState) *Meeting {
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

func GetState(meeting *Meeting) MeetingState {
	return meeting.state
}
