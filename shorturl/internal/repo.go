package internal

import (
	"github.com/nrednav/cuid2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repo interface {
	Create(url string) (ShortURL, error)
	Get(id string) (ShortURL, error)
}

type repo struct {
	db *gorm.DB
}

func newSqliteRepo() (*repo, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&ShortURL{})
	return &repo{db: db}, err
}

func (r *repo) Create(url string) (ShortURL, error) {
	shorturl := ShortURL{
		ID:     cuid2.Generate(),
		Target: url,
	}
	result := r.db.Create(&shorturl)
	return shorturl, result.Error
}

func (r *repo) Get(id string) (ShortURL, error) {
	var shorturl ShortURL
	result := r.db.Where("id = ?", id).First(&shorturl)
	return shorturl, result.Error
}
