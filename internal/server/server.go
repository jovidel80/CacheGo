package server

import (
	"fmt"
	"net/http"

	"github.com/jovidel80/cacheGo/internal/database"
)

type CacheServer struct {
	Database database.CacheDatabase
}

func (p *CacheServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cacheKey := r.Header.Get("X-Cache-Key")
	fmt.Fprint(w, p.Database.GetCacheByKey(cacheKey))
}

func GetCacheByKey(key string) string {
	if key == "123" {
		return "cache1"
	}

	if key == "456" {
		return "cache2"
	}

	return ""
}
