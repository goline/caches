package caches

import "time"

type CacheItem interface {
	// IsHit returns true if cache is hit, and false if miss
	IsHit() bool

	// Value allows to get an object from cache
	Value(v interface{}) error

	// WithValue sets an object into cache
	WithValue(v interface{}) CacheItem

	// Expire sets expiration for item
	Expire(t time.Duration) CacheItem
}

type CacheItemGetter interface {
	// String returns value as string
	String() (string, error)

	// Int returns value as int64
	Int() (int64, error)

	// Float returns value as float64
	Float() (float64, error)
}
