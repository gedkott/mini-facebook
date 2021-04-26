package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ProfileDatabaseModel struct {
	gorm.Model
	Name string
}

type ProfileDatabase struct {
	db *gorm.DB
}

func (pr *ProfileDatabase) GetAll() []ProfileDatabaseModel {
	var profiles []ProfileDatabaseModel
	result := pr.db.Find(&profiles)
	if result == nil || result.Error != nil {
		panic("failed to get all profiles")
	}

	return profiles
}

func (pr *ProfileDatabase) Create(p ProfileDatabaseModel) []ProfileDatabaseModel {
	result := pr.db.Create(&ProfileDatabaseModel{Name: p.Name})

	if result == nil || result.Error != nil {
		panic("failed to create profile")
	}

	return pr.GetAll()
}

func newProfileDatabase() ProfileDatabase {
	db, err := gorm.Open(sqlite.Open("test-profiles.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect to database")
	}

	// Migrate the schema
	db.AutoMigrate(&ProfileDatabaseModel{})
	return ProfileDatabase{
		db: db,
	}
}
