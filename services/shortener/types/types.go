package types

import "context"

type UrlDto struct {
	Id                int64
	ShortUrl, LongUrl string
}

type UrlStore interface {
	Create(context.Context, *UrlDto) error
	GetLongUrl(context.Context, string) (string, error)
	GetShortUrl(context.Context, string) (string, error)
}

type ShortenerService interface {
	ShortenUrl(context.Context, string) (string, error)
	GetLongUrl(context.Context, string) (string, error)
}
