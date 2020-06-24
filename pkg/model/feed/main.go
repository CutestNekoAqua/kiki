package feed

import (
	"gitea.code-infection.com/efertone/kiki/pkg/database"
	"gitea.code-infection.com/efertone/kiki/pkg/model"
)

// Add a new Feed to an Account.
func Add(name, user, url, provider string) error {
	db := database.NewDatabase()
	defer db.Close()

	var count int

	db.Connection().Model(&model.Account{}).Where("name = ?", user).Count(&count)

	if count < 1 {
		return UserDoesNotExistError{
			Name: user,
		}
	}

	db.Connection().Create(&model.Feed{
		Name:     name,
		User:     user,
		URL:      url,
		Provider: provider,
	})

	return nil
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
