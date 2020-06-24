package account

import (
	"gitea.code-infection.com/efertone/kiki/pkg/database"
	"gitea.code-infection.com/efertone/kiki/pkg/model"
)

// All returns all available Accounts.
func All() []*model.Account {
	db := database.NewDatabase()
	defer db.Close()

	var accounts []*model.Account

	db.Connection().Find(&accounts)

	return accounts
}

// Add a new Account.
func Add(name, token, publisher, baseURL string) error {
	db := database.NewDatabase()
	defer db.Close()

	var count int

	query := db.Connection().Model(&model.Account{}).Where("name = ?", name)
	query.Count(&count)

	if count > 0 {
		var account model.Account

		query.First(&account)

		return UserAlreadyExistsError{Account: account}
	}

	db.Connection().Create(&model.Account{
		Name:      name,
		APIToken:  token,
		Publisher: publisher,
		BaseURL:   baseURL,
	})

	return nil
}
