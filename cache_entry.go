package cacheManager

import "time"

type CacheEntry struct {
	Treshold time.Time
	Value    any
}
