package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/vhall1/shorturl/services/shortener/types"
)

const tableName = "url"

type UrlStore struct {
	db *sql.DB
}

func NewUrlStore(db *sql.DB) *UrlStore {
	return &UrlStore{db: db}
}

func (s *UrlStore) Create(ctx context.Context, urlDto *types.UrlDto) error {
	q := fmt.Sprintf(`
		INSERT INTO "%s" ("id", "shortUrl", "longUrl")
		VALUES ($1, $2, $3)
	`, tableName)

	_, err := s.db.ExecContext(ctx, q, urlDto.Id, urlDto.ShortUrl, urlDto.LongUrl)
	if err != nil {
		return err
	}

	return nil
}

func (s *UrlStore) FindByLongUrl(ctx context.Context, longUrl string) (string, error) {
	q := fmt.Sprintf(`SELECT "shortUrl" FROM "%s" WHERE "longUrl" = $1`, tableName)
	row := s.db.QueryRowContext(ctx, q, longUrl)

	var shortUrl string
	if err := row.Scan(&shortUrl); err != nil {
		return "", err
	}

	return shortUrl, nil
}

func (s *UrlStore) FindByShortUrl(ctx context.Context, shortUrl string) (string, error) {
	q := fmt.Sprintf(`SELECT "longUrl" FROM "%s" WHERE "shortUrl" = $1`, tableName)
	row := s.db.QueryRowContext(ctx, q, shortUrl)

	var longUrl string
	if err := row.Scan(&longUrl); err != nil {
		return "", err
	}

	return longUrl, nil
}
