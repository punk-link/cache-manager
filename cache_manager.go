package cacheManager

import "time"

type CacheManager interface {
	Remove(key string)
	Set(key string, value any, interval time.Duration)
	TryGet(key string) (any, bool)
}
