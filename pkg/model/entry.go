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
	// MaximumContentLength is the upper limit on one post.
	MaximumContentLength = 500
	// MinimumContentLength is the lower limit on one post.
	MinimumContentLength = 50
)

// Excerpt generates an except for an entry.
func (e Entry) Excerpt() string {
	if len(e.Content) < MaximumContentLength {
		return e.Content
	}

	// Find divider if exists
	index := strings.Index(e.Content[:MaximumContentLength], "----")

	// Last punctuation
	if index < MinimumContentLength {
		index = strings.LastIndexAny(e.Content[:MaximumContentLength], ".!?")
	}
	// Last newline
	if index < MinimumContentLength {
		index = strings.LastIndexAny(e.Content[:MaximumContentLength], "\n")
	}
	// Last whitespace
	if index < MinimumContentLength {
		index = strings.LastIndexAny(e.Content[:MaximumContentLength], " ")
	}
	// Hard cut
	if index < MinimumContentLength {
		index = MaximumContentLength
	}

	content := e.Content[:index+1]
	content = strings.Trim(content, "\n \r-")

	return content
}
