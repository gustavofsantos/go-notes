package core

import (
	"testing"
)

const TEXT = "A title"
const HOUR = "10"
const MINUTES = "30"
const STATE = TODO

func TestNewMeeting(t *testing.T) {
	meeting := NewMeeting(TEXT, HOUR, MINUTES, STATE)

	if meeting.text != TEXT {
		t.Fatalf("Expected meeting.text to be %s but received %s", TEXT, meeting.text)
	}

	if meeting.hour != HOUR {
		t.Fatalf("Expected meeting.hour to be %s but received %s", HOUR, meeting.hour)
	}

	if meeting.minutes != MINUTES {
		t.Fatalf("Expected meeting.minutes to be %s but received %s", MINUTES, meeting.minutes)
	}

	if meeting.state != STATE {
		t.Fatalf("Expected meeting.state to be %s but received %s", STATE, meeting.state)
	}
}

func TestGetText(t *testing.T) {
	meeting := NewMeeting(TEXT, HOUR, MINUTES, STATE)

	if GetText(meeting) != TEXT {
		t.Fatalf("Expected %s, got %s", TEXT, GetText(meeting))
	}
}

func TestParseLineLookingForMeeting(t *testing.T) {
	line := "- [ ] 10:00 A meeting"
	found, meeting := ParseLineLookingForMeeting(line)

	if found != true {
		t.Fatalf("Expected true, got %t", found)
	}

	if GetText(meeting) != "A meeting" {
		t.Fatalf("Expected 'A meeting', got %s", GetText(meeting))
	}

	if GetHour(meeting) != "10" {
		t.Fatalf("Expected '10', got %s", GetHour(meeting))
	}

	if GetMinutes(meeting) != "00" {
		t.Fatalf("Expected '00', got %s", GetMinutes(meeting))
	}

	if GetState(meeting) != "TODO" {
		t.Fatalf("Expected 'TODO', got %s", GetState(meeting))
	}
}

func TestParseMeetings(t *testing.T) {
	journal := `
- [ ] A task
- [ ] 09:00 A meeting
`
	meetings := ParseMeetings(journal)

	if len(meetings) != 1 {
		t.Fatalf("Expected len of meetings to be 1, got %d", len(meetings))
	}
}
