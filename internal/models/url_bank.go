package models

import (
	"database/sql"
	"errors"
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
	var urlBank UrlBank
	query := `SELECT id, actual_url, short_url, total_hit, created_at FROM url_bank WHERE short_url = $1`
	err := m.DB.QueryRow(query, shortUrl).Scan(
		&urlBank.ID,
		&urlBank.ActualUrl,
		&urlBank.ShortUrl,
		&urlBank.TotalHit,
		&urlBank.CreatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("models: no matching record found")
		default:
			return nil, err
		}
	}
	return &urlBank, nil
}

func (m *UrlBankModel) UpdateHitCountAndGet(urlBank *UrlBank) error {
	query := `UPDATE url_bank SET total_hit = total_hit + 1 WHERE short_url = $1 
	RETURNING id, actual_url, short_url, total_hit, created_at`
	args := []interface{}{
		urlBank.ShortUrl,
	}
	return m.DB.QueryRow(query, args...).Scan(
		&urlBank.ID,
		&urlBank.ActualUrl,
		&urlBank.ShortUrl,
		&urlBank.TotalHit,
		&urlBank.CreatedAt,
	)
}

func (m *UrlBankModel) AllUrl() ([]UrlBank, error) {

	return nil, nil
}
