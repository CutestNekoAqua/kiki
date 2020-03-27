package database

import (
	"log"

	"gitea.code-infection.com/efertone/kiki/pkg/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	connectionString string
)

type Database struct {
	connection *gorm.DB
}

func (d *Database) Close() {
	d.connection.Close()
}

func (d *Database) Migrate() {
	d.Connection().AutoMigrate(
		&model.Account{},
		&model.Feed{},
		&model.Entry{},
	)
	d.Close()
}

func (d *Database) Connection() *gorm.DB {
	if d.connection == nil {
		var err error
		d.connection, err = gorm.Open("postgres", connectionString)

		if err != nil {
			log.Fatalf("Database connection error: %s\n", err)
		}
	}

	return d.connection
}

func NewDatabase() *Database {
	return &Database{}
}
