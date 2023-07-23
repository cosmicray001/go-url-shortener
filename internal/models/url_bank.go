package models

import (
	"database/sql"
	"time"
)

type UrlBank struct {
	ID        int64     `json:"id"`
	ActualUrl string    `json:"actual_url"`
	ShortUrl  string    `json:"short_url"`
	TotalHit  int64     `json:"total_hit"`
	CreatedAt time.Time `json:"created_at"`
}

type UrlBankModel struct {
	DB *sql.DB
}

func (m *UrlBankModel) CheckExistUrl(shortUrl string) (bool, error) {
	var exists bool

	query := "SELECT EXISTS(SELECT true FROM url_bank WHERE short_url = $1)"

	err := m.DB.QueryRow(query, shortUrl).Scan(&exists)

	return exists, err
}

func (m *UrlBankModel) Insert(urlBank *UrlBank) error {
	query := `INSERT INTO url_bank (actual_url, short_url, total_hit, created_at) VALUES($1, $2, $3, $4) RETURNING id, created_at`
	args := []interface{}{urlBank.ActualUrl, urlBank.ShortUrl, 0, time.Now().UTC()}
	err := m.DB.QueryRow(query, args...).Scan(&urlBank.ID, &urlBank.CreatedAt)
	return err
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
