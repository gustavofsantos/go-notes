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
