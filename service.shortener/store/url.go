package store

import (
	"context"
	"database/sql"

	"github.com/vhall1/shorturl/service.shortener/common"
)

type UrlStore struct {
	db *sql.DB
}

func NewUrlStore(db *sql.DB) *UrlStore {
	return &UrlStore{db: db}
}

func (s *UrlStore) Create(ctx context.Context, url *common.Url) error {
	q := `
		INSERT INTO url (id, shortUrl, longUrl)
		VALUES (?, ?, ?)
	`

	_, err := s.db.ExecContext(ctx, q, url.Id, url.ShortUrl, url.LongUrl)

	return err
}
