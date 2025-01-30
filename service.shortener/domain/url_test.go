package domain

import (
	"context"
	"testing"
)

func TestUrlService(t *testing.T) {
	s := NewUrlService()

	shortUrl, err := s.ShortenUrl(context.Background(), "example.com")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	got := shortUrl
	want := "14q60P"
	if got != want {
		t.Fatalf("expected shortUrl %v, got %v", want, got)
	}
}
