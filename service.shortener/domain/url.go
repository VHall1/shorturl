package domain

import (
	"context"
	"database/sql"
	"errors"

	"github.com/vhall1/shorturl/service.shortener/common"
	"github.com/vhall1/shorturl/service.shortener/store"
)

type UrlService struct {
	urlStore *store.UrlStore

	snowflakeService *SnowflakeService
}

func NewUrlService(urlStore *store.UrlStore, snowflakeService *SnowflakeService) *UrlService {
	return &UrlService{
		urlStore:         urlStore,
		snowflakeService: snowflakeService,
	}
}

func (s *UrlService) ShortenUrl(ctx context.Context, longUrl string) (string, error) {
	shortUrl, err := s.urlStore.FindByLongUrl(ctx, longUrl)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", err
	} else if shortUrl != "" {
		return shortUrl, nil
	}

	id, err := s.snowflakeService.Generate(ctx)
	if err != nil {
		return "", err
	}

	err = s.urlStore.Create(ctx, &common.Url{
		Id:       id,
		ShortUrl: common.Base62(id),
		LongUrl:  longUrl,
	})

	if err != nil {
		return "", err
	}

	return shortUrl, nil
}

func (s *UrlService) GetRedirectUrl(ctx context.Context, shortUrl string) (string, error) {
	longUrl, err := s.urlStore.FindByShortUrl(ctx, shortUrl)

	if err != nil {
		return "", err
	}

	return longUrl, nil
}
