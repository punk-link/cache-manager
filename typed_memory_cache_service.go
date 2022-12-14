package cacheManager

import (
	"time"
)

type MemoryCacheService[T any] struct {
	cacheService *InnerMemoryCacheService
}

func New[T any]() CacheManager[T] {
	return &MemoryCacheService[T]{
		cacheService: GetInnerCache(),
	}
}

func (t *MemoryCacheService[T]) Remove(key string) {
	if key == "" {
		return
	}

	t.Remove(key)
}

func (t *MemoryCacheService[T]) Set(key string, value T, interval time.Duration) {
	if key == "" {
		return
	}

	t.cacheService.Set(key, value, interval)
}

func (t *MemoryCacheService[T]) TryGet(key string) (T, bool) {
	var emptyOfT T
	if key == "" {
		return emptyOfT, false
	}

	entry, isCached := t.cacheService.TryGet(key)
	if isCached {
		return entry.(T), true
	}

	return emptyOfT, false
}
