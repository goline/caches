package caches

import "time"

type Cache interface {
	// Get returns an appropriate cache item
	Get(id string) CacheItem

	// Set adds values into cache
	Set(id string, value interface{}, expiration time.Duration) Cache

	// Put puts item into cache
	Put(item CacheItem) Cache
}

// New returns a cache instance by connection url
// Example:
// 		caches.New("redis://user:password@localhost:6379/0")
// It will return Redis cache instance with options:
// 		user=user password=password
//		host=localhost port=6379
//		database=0 (default database)
func New(url string) Cache {
	return nil
}
