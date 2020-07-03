package hashtag

import (
	"gitea.code-infection.com/efertone/kiki/pkg/database"
	"gitea.code-infection.com/efertone/kiki/pkg/model"
)

// Add a new Feed to an Account.
func Add(feedID uint, value string) error {
	db := database.NewDatabase()
	defer db.Close()

	var count int

	db.Connection().Model(&model.Feed{}).Where("ID = ?", feedID).Count(&count)

	if count < 1 {
		return FeedDoesNotExistError{
			ID: feedID,
		}
	}

	db.Connection().Create(&model.HashTag{
		FeedID: feedID,
		Value:  value,
	})

	return nil
}
