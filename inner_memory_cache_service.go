package cacheManager

import (
	"sync"
	"time"
)

type InnerMemoryCacheService struct {
	cache map[string]CacheEntry
	mutex sync.Mutex
}

func GetInnerCache() *InnerMemoryCacheService {
	if _cache != nil {
		return _cache
	}

	cache := make(map[string]CacheEntry)
	service := InnerMemoryCacheService{
		cache: cache,
	}

	go service.watch()

	_cache = &service
	return _cache
}

func (t *InnerMemoryCacheService) Remove(key string) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	delete(t.cache, key)
}

func (t *InnerMemoryCacheService) Set(key string, value any, interval time.Duration) {
	treshold := time.Now().UTC().Add(interval)
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.cache[key] = CacheEntry{
		Treshold: treshold,
		Value:    value,
	}
}

func (t *InnerMemoryCacheService) TryGet(key string) (any, bool) {
	if entry, isCached := t.cache[key]; isCached {
		return entry.Value, true
	}

	return nil, false
}

func (t *InnerMemoryCacheService) watch() {
	ticker := time.NewTicker(LIFETIME_VALIDATION_INTERVAL)
	for range ticker.C {
		for key, value := range t.cache {
			if time.Now().UTC().After(value.Treshold) {
				t.Remove(key)
			}
		}
	}
}

var _cache *InnerMemoryCacheService

const LIFETIME_VALIDATION_INTERVAL = 5 * time.Second
