package domain

import (
	"context"

	"github.com/vhall1/shorturl/service.shortener/common"
	"github.com/vhall1/shorturl/service.shortener/store"
)

type UrlService struct {
	urlStore *store.UrlStore
}

func NewUrlService(urlStore *store.UrlStore) *UrlService {
	return &UrlService{urlStore: urlStore}
}

// TODO: Implement this
func (s *UrlService) ShortenUrl(ctx context.Context, longUrl string) (string, error) {
	shortUrl := "14q60P"

	err := s.urlStore.Create(ctx, &common.Url{
		Id:       1,
		ShortUrl: shortUrl,
		LongUrl:  longUrl,
	})

	if err != nil {
		return "", err
	}

	return shortUrl, nil
}
