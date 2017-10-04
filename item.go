package caches

import "time"

type CacheItem interface {
	// IsHit returns true if cache is hit, and false if miss
	IsHit() bool

	// Value gets value into input argument
	Value(v interface{}) error

	// WithValue sets input value into cache
	WithValue(v interface{}) CacheItem

	// Expire sets expiration for item
	Expire(t time.Duration) CacheItem
}
