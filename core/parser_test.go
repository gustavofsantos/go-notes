package core

import (
	"testing"
)

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

	if GetState(meeting) != TODO {
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
