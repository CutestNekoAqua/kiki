package model

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Entry database model
type Entry struct {
	*gorm.Model
	FeedID   uint
	Link     string
	Title    string
	EntryID  string
	Content  string
	PostedAt *time.Time
}

// Excerpt generates an except for an entry
func (e Entry) Excerpt() string {
	if len(e.Content) < 500 {
		return e.Content
	}

	var index int
	index = strings.Index(e.Content[:500], "----")
	if index < 50 {
		index = strings.LastIndexAny(e.Content[:500], ".!?")
	}
	content := e.Content[:index+1]
	content = strings.Trim(content, "\n \r-")
	return content
}
