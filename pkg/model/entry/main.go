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
