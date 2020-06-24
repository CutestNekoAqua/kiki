package entry

import (
	"time"

	"gitea.code-infection.com/efertone/kiki/pkg/database"
	"gitea.code-infection.com/efertone/kiki/pkg/model"
)

// MarkAsPosted marks an Entry as Posted.
func MarkAsPosted(entry *model.Entry) {
	now := time.Now()
	entry.PostedAt = &now
	db := database.NewDatabase()

	defer db.Close()
	db.Connection().Save(entry)
}

// NextPending returns the next pending Entry for a Feed.
func NextPending(feed *model.Feed) *model.Entry {
	db := database.NewDatabase()
	defer db.Close()

	var entry model.Entry

	db.Connection().Where("feed_id = ? and posted_at is null", feed.ID).First(&entry)

	if entry.ID == 0 {
		return nil
	}

	return &entry
}

// Exists checks if a given Entry already exists in the database or not.
func Exists(entry *model.Entry) bool {
	var count int

	db := database.NewDatabase()
	defer db.Close()

	db.Connection().
		Model(&model.Entry{}).
		Where(&model.Entry{FeedID: entry.FeedID, EntryID: entry.EntryID}).
		Or(&model.Entry{FeedID: entry.FeedID, Link: entry.Link}).
		Count(&count)

	return count > 0
}
