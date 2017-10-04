package caches

type Cache interface {
	// Get returns an appropriate cache item
	Get(id string) CacheItem

	// Set puts item into cache
	Set(item CacheItem) Cache
}
