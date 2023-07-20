package models

import (
	"database/sql"
	"time"
)

type UrlBank struct {
	ID         int64
	ActualUrl  string
	ShortUrl   string
	TotalHit   int64
	CreatedAt  time.Time
	ModifiedAt time.Time
}

type UrlBankModel struct {
	DB *sql.DB
}

func (m *UrlBankModel) CheckExistUrl(actualUrl string) (bool, error) {
	return false, nil
}

func (m *UrlBankModel) Insert(actualUrl string, shortUrl string) error {
	return nil
}

func (m *UrlBankModel) Get(shortUrl string) (*UrlBank, error) {
	return nil, nil
}

func (m *UrlBankModel) UpdateHitCount(ID int64, value int64) error {
	return nil
}
