package account

import (
	"log"

	"gitea.code-infection.com/efertone/kiki/pkg/database"
	"gitea.code-infection.com/efertone/kiki/pkg/model"
)

func All() []*model.Account {
	db := database.NewDatabase()
	defer db.Close()

	var accounts []*model.Account
	db.Connection().Find(&accounts)

	return accounts
}

func Add(name, token string) {
	db := database.NewDatabase()
	defer db.Close()

	var count int
	db.Connection().Find(&model.Account{Name: name}).Count(&count)

	if count > 1 {
		log.Fatalln("User already exists")
	}

	db.Connection().Create(&model.Account{Name: name, APIToken: token})

}

func init() {}
