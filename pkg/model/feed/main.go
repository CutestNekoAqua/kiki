package feed

import (
	"log"

	"gitea.code-infection.com/efertone/kiki/pkg/database"
	"gitea.code-infection.com/efertone/kiki/pkg/model"
)

// Add a new Feed to an Account.
func Add(name, user, url, provider string) {
	db := database.NewDatabase()
	defer db.Close()

	var count int

	db.Connection().Find(&model.Account{Name: user}).Count(&count)

	if count < 1 {
		log.Fatalln("User does not exist")
	}

	db.Connection().Create(&model.Feed{
		Name:     name,
		User:     user,
		URL:      url,
		Provider: provider,
	})
}

// All returns all available Feeds.
func All() []*model.Feed {
	db := database.NewDatabase()
	defer db.Close()

	var feeds []*model.Feed

	db.Connection().Find(&feeds)

	return feeds
}

// AllFor returns all Feed for an Account.
func AllFor(acc *model.Account) []*model.Feed {
	db := database.NewDatabase()
	defer db.Close()

	var feeds []*model.Feed

	db.Connection().Where(&model.Feed{User: acc.Name}).Find(&feeds)

	return feeds
}
