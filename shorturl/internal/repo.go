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
	id func() string
}

func newSqliteRepo() (*repo, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&ShortURL{})
	if err != nil {
		return nil, err
	}

	id, err := cuid2.Init(cuid2.WithLength(6))
	if err != nil {
		return nil, err
	}
	return &repo{db: db, id: id}, err
}

func (r *repo) Create(url string) (ShortURL, error) {
	shorturl := ShortURL{
		ID:     r.id(),
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
