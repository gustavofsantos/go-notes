package core

import (
	"testing"
	"time"
)

func TestGetTodayDateAsString(t *testing.T) {
	now := time.UnixMilli(1673189208743)
	todayFile := GetTodayDateAsString(now)

	if todayFile != "2023-01-08.md" {
		t.Fatalf("Expected 2023-01-08.md, got %s", todayFile)
	}
}

func TestGetFilePath(t *testing.T) {
	path := "/home/user/notes/"
	filename := "2023-01-08.md"
	filepath := GetFilePath(path, filename)

	if filepath != "/home/user/notes/journal/2023-01-08.md" {
		t.Fatalf("Expected /home/user/notes/journal/2023-01-08.md, got %s", filepath)
	}
}
