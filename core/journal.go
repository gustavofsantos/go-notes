package core

import (
	"fmt"
	"path"
	"time"
)

func GetTodayDateAsString(now time.Time) string {
	layout := "2006-01-02"
	return fmt.Sprintf("%s.md", now.Format(layout))
}

func GetFilePath(notesPath string, filename string) string {
	return path.Join(notesPath, "/journal/", filename)
}
