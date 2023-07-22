package models

import (
	"database/sql"
	"time"
)

type UrlBank struct {
	ID        int64
	ActualUrl string
	ShortUrl  string
	TotalHit  int64
	CreatedAt time.Time
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
	query := `SELECT id, actual_url, short_url, created_at FROM url_bank WHERE id = $1`
	var urlBank UrlBank
	err := m.DB.QueryRow(query, 1).Scan(&urlBank.ID, &urlBank.ActualUrl, &urlBank.ShortUrl, &urlBank.CreatedAt)
	_ = err
	return nil
}

func (m *UrlBankModel) AllUrl() ([]UrlBank, error) {

	return nil, nil
}
