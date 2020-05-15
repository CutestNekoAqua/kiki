package model

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Entry database model.
type Entry struct {
	*gorm.Model
	FeedID   uint
	Link     string
	Title    string
	EntryID  string
	Content  string
	PostedAt *time.Time
}

const (
	MaximumContentLength = 500
	MinimumContentLength = 50
)

// Excerpt generates an except for an entry.
func (e Entry) Excerpt() string {
	if len(e.Content) < MaximumContentLength {
		return e.Content
	}

	index := strings.Index(e.Content[:MaximumContentLength], "----")
	if index < MinimumContentLength {
		index = strings.LastIndexAny(e.Content[:MaximumContentLength], ".!?")
	}

	content := e.Content[:index+1]
	content = strings.Trim(content, "\n \r-")

	return content
}
