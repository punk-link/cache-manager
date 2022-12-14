package cacheManager

import "time"

type CacheManager[T any] interface {
	Remove(key string)
	Set(key string, value T, interval time.Duration)
	TryGet(key string) (T, bool)
}
