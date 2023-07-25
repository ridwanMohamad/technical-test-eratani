package seeds

import (
	"gorm.io/gorm"
)

func isTableEmpty(db *gorm.DB, table interface{}) bool {
	var count int64
	db.Model(table).Count(&count)
	return count == 0
}

type seeds struct {
	db   *gorm.DB
	json map[string]string
}

func InitialSeeds(db *gorm.DB, config map[string]string) Seeds {
	return &seeds{
		db:   db,
		json: config,
	}
}

type Seeds interface {
	SeedUsers() error
	SeedBelanja() error
}
