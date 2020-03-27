package feed

import (
	"log"
	"time"

	"gitea.code-infection.com/efertone/kiki/pkg/database"
	"gitea.code-infection.com/efertone/kiki/pkg/model"
)

func AllFor(acc *model.Account) []*model.Feed {
	db := database.NewDatabase()
	defer db.Close()

	var feeds []*model.Feed
	db.Connection().Where(&model.Feed{User: acc.Name}).Find(&feeds)

	return feeds
}

func NextPendingEntries(feed *model.Feed) *model.Entry {
	db := database.NewDatabase()
	defer db.Close()

	var entry model.Entry
	db.Connection().Where("feed_id = ? and posted_at is null", feed.ID).First(&entry)

	if entry.ID == 0 {
		return nil
	}

	return &entry
}

func All() []*model.Feed {
	db := database.NewDatabase()
	defer db.Close()

	var feeds []*model.Feed
	db.Connection().Find(&feeds)

	return feeds
}

func Add(name, user, url string) {
	db := database.NewDatabase()
	defer db.Close()

	var count int
	db.Connection().Find(&model.Account{Name: user}).Count(&count)

	if count < 1 {
		log.Fatalln("User does not exist")
	}

	db.Connection().Create(&model.Feed{Name: name, User: user, URL: url})
}

func Get(feed *model.Feed) {
	entries, err := Download(feed)
	if err != nil {
		log.Printf("[%s] Error: %s\n", feed.Name, err)
		return
	}

	db := database.NewDatabase()
	defer db.Close()

	for i := len(entries) - 1; i >= 0; i-- {
		entry := entries[i]
		var count int
		db.Connection().
			Model(&model.Entry{}).
			Where(&model.Entry{FeedID: entry.FeedID, EntryID: entry.EntryID}).
			Or(&model.Entry{FeedID: entry.FeedID, Link: entry.Link}).
			Count(&count)
		if count > 0 {
			continue
		}
		db.Connection().Create(entry)
	}
}

func MarAsPosted(entry *model.Entry) {
	now := time.Now()
	entry.PostedAt = &now
	db := database.NewDatabase()
	defer db.Close()
	db.Connection().Save(entry)
}
