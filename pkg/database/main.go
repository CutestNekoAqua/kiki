package database

import (
	"log"

	"gitea.code-infection.com/efertone/kiki/pkg/model"
	"github.com/jinzhu/gorm"

	// Required by gorm to access postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	connectionString string //nolint:gochecknoglobals
)

// Database is a wrapper for gorm.DB.
type Database struct {
	connection *gorm.DB
}

// Close is a wrapper to close the Database Connection.
func (d *Database) Close() {
	if d.connection != nil {
		d.connection.Close()
	}
}

// Migrate is a wrapper to call AutoMigrate on a Database Connection.
func (d *Database) Migrate() {
	d.Connection().AutoMigrate(
		&model.Account{},
		&model.Feed{},
		&model.Entry{},
		&model.HashTag{},
	)
	d.Close()
}

// Connection creates a new database connection or if exists returns with the existing one.
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

// NewDatabase create a new Database wrapper.
func NewDatabase() *Database {
	return &Database{}
}
