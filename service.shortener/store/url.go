package store

import (
	"context"
	"database/sql"
	"errors"

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

func (s *UrlStore) FindByLongUrl(ctx context.Context, longUrl string) (string, error) {
	row := s.db.QueryRowContext(ctx, "SELECT * FROM url WHERE longUrl = ?", longUrl)

	url, err := scanIntoUrl(row)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", err
	}

	return url.ShortUrl, nil
}

func (s *UrlStore) FindByShortUrl(ctx context.Context, shortUrl string) (string, error) {
	row := s.db.QueryRowContext(ctx, "SELECT * FROM url WHERE shortUrl = ?", shortUrl)

	url, err := scanIntoUrl(row)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", err
	}

	return url.LongUrl, nil
}

func scanIntoUrl(row *sql.Row) (*common.Url, error) {
	url := new(common.Url)

	if err := row.Scan(&url.Id, &url.LongUrl, &url.ShortUrl); err != nil {
		return nil, err
	}

	return url, nil
}
