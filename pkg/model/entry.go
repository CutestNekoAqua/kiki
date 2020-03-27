package model

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Entry struct {
	*gorm.Model
	FeedID   uint
	Link     string
	Title    string
	EntryID  string
	Content  string
	PostedAt *time.Time
}

func (e Entry) Excerpt() string {
	if len(e.Content) < 300 {
		return e.Content
	}
	index := strings.Index(e.Content, "\n\n")
	if index < 0 {
		index = strings.IndexAny(e.Content, ".!?")
	}
	return e.Content[:index+1]
}
