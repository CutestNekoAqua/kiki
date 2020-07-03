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

// FindByID looks up and returns with a Feed by its ID.
func FindByID(feedID uint) (*model.Feed, error) {
	db := database.NewDatabase()
	defer db.Close()

	var feed model.Feed

	db.Connection().Where("ID = ?", feedID).First(&feed)

	if feed.ID == 0 {
		return nil, DoesNotExistError{feedID}
	}

	return &feed, nil
}

// HashTagsFor return all associated HashTags for a Feed.
func HashTagsFor(f *model.Feed) []*model.HashTag {
	db := database.NewDatabase()
	defer db.Close()

	var hashTags []*model.HashTag

	db.Connection().Model(f).Related(&hashTags)

	return hashTags
}

// PrefixedHashTagValuesFor returns with all the HashTags with
// prefixed (#) array of strings.
func PrefixedHashTagValuesFor(f *model.Feed) []string {
	hashTags := HashTagsFor(f)
	tags := make([]string, 0)

	for _, tag := range hashTags {
		tags = append(tags, tag.PrefixedValue())
	}

	return tags
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
