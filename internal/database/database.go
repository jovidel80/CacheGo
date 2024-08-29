package database

type CacheDatabase interface {
	GetCacheByKey(key string) string
}
