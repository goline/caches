package caches

type CacheManager interface {
	// Tag returns cache instance for specific tag
	Tag(tag string) (Cache, bool)

	// Register adds a cache resource to manager
	Register(link string, tags ...string) CacheManager
}
