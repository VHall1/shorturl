package domain

import "context"

type Url struct {
	Id                int64
	ShortUrl, LongUrl string
}

type UrlService struct{}

func NewUrlService() *UrlService {
	return &UrlService{}
}

// TODO: Implement this
func (s *UrlService) ShortenUrl(ctx context.Context, longUrl string) (string, error) {
	return "14q60P", nil
}
