package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockSnowflake struct{}

const mockSnowflakeId int64 = 7281258113146028032

func (s *MockSnowflake) Generate() int64 {
	return mockSnowflakeId
}

func TestHandleGetSnowflake(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	h := NewSnowflakeHttpHandler(&MockSnowflake{})

	h.HandleGetSnowflake(h.snowflakeService).ServeHTTP(w, r)

	assertStatus(t, w.Code, http.StatusOK)
	assertHeader(t, w.Header().Get("Content-Type"), "application/json")

	var res map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&res); err != nil {
		t.Fatalf("failed decoding json: %v", err)
	}

	id, present := res["id"]
	if !present {
		t.Fatalf("response body does not contain key \"id\"")
	}

	// JSON unmarshals numbers as float64 by default
	idFloat, ok := id.(float64)
	if !ok {
		t.Fatalf("\"id\" is not a valid number, got: %v", idFloat)
	}

	id = int64(idFloat)
	if id != mockSnowflakeId {
		t.Fatalf("expected \"id\" to be %v, got %v", mockSnowflakeId, id)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status code, got %v, want %v", got, want)
	}
}

func assertHeader(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct header value, got %v, want %v", got, want)
	}
}
