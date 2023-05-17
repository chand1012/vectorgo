package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// create a new sqlite db
func NewDB(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// perform auto migration
	err = db.AutoMigrate(&Embedding{}, &EmbeddingValue{})

	return db, err
}
