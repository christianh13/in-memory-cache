package cache

import (
	"errors"

	"github.com/in-memory-cache/evictionmanager"
)

type InMemoryCache struct {
	cache           map[string]string
	Limit           int
	EvictionManager evictionmanager.EvictionManager
}

func NewInMemoryCache(limit int, evictionManager evictionmanager.EvictionManager) *InMemoryCache {
	if limit < 1 {
		panic(errors.New("limit_must_be_positive"))
	}
	return &InMemoryCache{cache: make(map[string]string), Limit: limit, EvictionManager: evictionManager}
}

func (inMemoryCache InMemoryCache) Add(key, value string) int {
	result := 0
	if _, exists := inMemoryCache.cache[key]; exists {
		inMemoryCache.cache[key] = value
		result = 1
	} else {
		if len(inMemoryCache.cache) == inMemoryCache.Limit {
			keyToBeDeleted := inMemoryCache.EvictionManager.HandleOverLimit()
			delete(inMemoryCache.cache, keyToBeDeleted)
		}
		inMemoryCache.EvictionManager.Push(key)
		inMemoryCache.cache[key] = value
	}
	return result
}

func (inMemoryCache InMemoryCache) Get(key string) string {
	if value, exists := inMemoryCache.cache[key]; exists {
		inMemoryCache.EvictionManager.Use(key)
		return value
	}
	return ""
}

func (inMemoryCache InMemoryCache) Clear() int {
	capacity := len(inMemoryCache.cache)
	inMemoryCache.EvictionManager.Clear()
	for k := range inMemoryCache.cache {
		delete(inMemoryCache.cache, k)
	}
	return capacity
}

func (inMemoryCache InMemoryCache) Keys() []string {
	keys := make([]string, 0, len(inMemoryCache.cache))
	for k := range inMemoryCache.cache {
		keys = append(keys, k)
	}
	return keys
}
