package domain

import (
	"context"
	"database/sql"
	"testing"

	"github.com/vhall1/shorturl/service.shortener/store"
)

func TestUrlService(t *testing.T) {
	str := store.NewUrlStore(&sql.DB{})
	snow := NewSnowflakeService("http://snowflake/")
	svc := NewUrlService(str, snow)

	shortUrl, err := svc.ShortenUrl(context.Background(), "example.com")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	got := shortUrl
	want := "14q60P"
	if got != want {
		t.Fatalf("expected shortUrl %v, got %v", want, got)
	}
}
