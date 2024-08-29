package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubCacheDatabase struct {
	cache map[string]string
}

func (s *StubCacheDatabase) GetCacheByKey(key string) string {
	return s.cache[key]
}

func TestGETCache(t *testing.T) {
	database := StubCacheDatabase{
		cache: map[string]string{
			"123": "cache1",
			"456": "cache2",
		},
	}

	server := CacheServer{&database}
	t.Run("returns cache when unique cache key: 123", func(t *testing.T) {
		request := newGetRequest("123")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "cache1"

		assertResponseBody(t, got, want)
	})

	t.Run("returns cache when unique cache key: 456", func(t *testing.T) {
		request := newGetRequest("456")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "cache2"

		assertResponseBody(t, got, want)
	})
}

func newGetRequest(cacheKey string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/cache", nil)
	request.Header.Add("X-Cache-Key", cacheKey)
	return request
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
