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

	_, err := s.db.Exec(q, urlDto.Id, urlDto.ShortUrl, urlDto.LongUrl)
	if err != nil {
		return err
	}

	return nil
}

func (s *UrlStore) GetLongUrl(ctx context.Context, shortUrl string) (string, error) {
	var longUrl string

	// tableName is a hardcoded const and doesn't come from user input, so should be safe
	// to do simple string interpolation here
	q := fmt.Sprintf(`SELECT "longUrl" FROM "%s" WHERE "shortUrl" = $1`, tableName)
	// TODO: actually do something with ctx
	row := s.db.QueryRow(q, shortUrl)

	if err := row.Scan(&longUrl); err != nil {
		return "", err
	}

	return longUrl, nil
}

func (s *UrlStore) GetShortUrl(ctx context.Context, longUrl string) (string, error) {
	var shortUrl string

	// tableName is a hardcoded const and doesn't come from user input, so should be safe
	// to do simple string interpolation here
	q := fmt.Sprintf(`SELECT "shortUrl" FROM "%s" WHERE "longUrl" = $1`, tableName)
	// TODO: actually do something with ctx
	row := s.db.QueryRow(q, longUrl)

	if err := row.Scan(&shortUrl); err != nil {
		return "", err
	}

	return shortUrl, nil
}
