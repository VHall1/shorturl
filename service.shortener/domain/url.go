package domain

import "context"

type URL struct {
	ID                int64
	ShortURL, LongURL string
}

type URLService struct{}

func NewURLService() *URLService {
	return &URLService{}
}

// TODO: Implement this
func (s *URLService) ShortenURL(ctx context.Context, longURL string) (string, error) {
	return "14q60P", nil
}
