package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/vhall1/shorturl/services/shortener/store"
	"github.com/vhall1/shorturl/services/shortener/types"
)

type ShortenerService struct {
	urlStore *store.UrlStore
}

func NewShortenerService(urlStore *store.UrlStore) *ShortenerService {
	return &ShortenerService{
		urlStore: urlStore,
	}
}

func (s *ShortenerService) ShortenUrl(ctx context.Context, longUrl string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	// do we have this url in the db already?
	shortUrl, err := s.urlStore.FindByLongUrl(ctx, longUrl)
	if err != nil {
		return "", err
	}

	// url already stored in the db, return that instead of storing a new copy
	if shortUrl != "" {
		return shortUrl, nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://snowflake/", nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	var jsonRes *struct{ Id int64 }
	err = json.NewDecoder(res.Body).Decode(&jsonRes)
	if err != nil {
		return "", err
	}

	shortUrl = Base62(jsonRes.Id)

	err = s.urlStore.Create(ctx, &types.UrlDto{
		Id:       jsonRes.Id,
		ShortUrl: shortUrl,
		LongUrl:  longUrl,
	})
	if err != nil {
		return "", err
	}

	return shortUrl, nil
}

func (s *ShortenerService) GetRedirectUrl(ctx context.Context, shortUrl string) (string, error) {
	longUrl, err := s.urlStore.FindByShortUrl(ctx, shortUrl)

	if err != nil {
		return "", err
	}

	return longUrl, nil
}
