package types

import "context"

type UrlDto struct {
	Id                int64
	ShortUrl, LongUrl string
}

type UrlStore interface {
	Create(context.Context, *UrlDto) error
	FindByLongUrl(context.Context, string) (string, error)
	FindByShortUrl(context.Context, string) (string, error)
}

type ShortenerService interface {
	ShortenUrl(context.Context, string) (string, error)
	GetRedirectUrl(context.Context, string) (string, error)
}
