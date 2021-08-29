package evictionmanager

import (
	"testing"

	"github.com/in-memory-cache/cache"
	"github.com/in-memory-cache/evictionmanager"
	"github.com/stretchr/testify/assert"
)

func TestLRUNormalCase(t *testing.T) {
	lruEvictionManager := evictionmanager.NewLRUEvictionManager()
	evictionCache := cache.NewInMemoryCache(3, lruEvictionManager)
	assert.Equal(t, 0, evictionCache.Add("key1", "value1"))
	assert.Equal(t, 0, evictionCache.Add("key2", "value2"))
	assert.Equal(t, 0, evictionCache.Add("key3", "value3"))
	assert.Equal(t, 1, evictionCache.Add("key2", "value2.1"))
	assert.Equal(t, "value3", evictionCache.Get("key3"))
	assert.Equal(t, "value1", evictionCache.Get("key1"))
	assert.Equal(t, "value2.1", evictionCache.Get("key2"))
	assert.ElementsMatch(t, []string{"key1", "key2", "key3"}, evictionCache.Keys())
	assert.Equal(t, 0, evictionCache.Add("key4", "value4"))
	assert.ElementsMatch(t, []string{"key1", "key2", "key4"}, evictionCache.Keys())
	assert.Equal(t, 3, evictionCache.Clear())
	assert.Equal(t, []string{}, evictionCache.Keys())
}

func TestLRULimit1(t *testing.T) {
	lfuEvictionManager := evictionmanager.NewLFUEvictionManager()
	evictionCache := cache.NewInMemoryCache(1, lfuEvictionManager)
	assert.Equal(t, 0, evictionCache.Add("key1", "value1"))
	assert.Equal(t, 0, evictionCache.Add("key2", "value2"))
	assert.Equal(t, 0, evictionCache.Add("key3", "value3"))
	assert.Equal(t, 0, evictionCache.Add("key2", "value2.1"))
	assert.Equal(t, "", evictionCache.Get("key3"))
	assert.Equal(t, "", evictionCache.Get("key1"))
	assert.Equal(t, "value2.1", evictionCache.Get("key2"))
	assert.ElementsMatch(t, []string{"key2"}, evictionCache.Keys())
	assert.Equal(t, 0, evictionCache.Add("key4", "value4"))
	assert.Equal(t, 0, evictionCache.Add("key1", "value1.1"))
	assert.ElementsMatch(t, []string{"key1"}, evictionCache.Keys())
	assert.Equal(t, 1, evictionCache.Clear())
	assert.ElementsMatch(t, []string{}, evictionCache.Keys())
}

func TestLRUCacheNotUsedYet(t *testing.T) {
	lfuEvictionManager := evictionmanager.NewLFUEvictionManager()
	evictionCache := cache.NewInMemoryCache(3, lfuEvictionManager)
	assert.Equal(t, 0, evictionCache.Add("key1", "value1"))
	assert.Equal(t, 0, evictionCache.Add("key2", "value2"))
	assert.Equal(t, 0, evictionCache.Add("key3", "value3"))
	assert.ElementsMatch(t, []string{"key1", "key2", "key3"}, evictionCache.Keys())
	assert.Equal(t, 0, evictionCache.Add("key4", "value4"))
	assert.ElementsMatch(t, []string{"key2", "key3", "key4"}, evictionCache.Keys())
}

func TestLRUEvictedKeyAddedAgain(t *testing.T) {
	lfuEvictionManager := evictionmanager.NewLFUEvictionManager()
	evictionCache := cache.NewInMemoryCache(3, lfuEvictionManager)
	assert.Equal(t, 0, evictionCache.Add("key1", "value1"))
	assert.Equal(t, 0, evictionCache.Add("key2", "value2"))
	assert.Equal(t, 0, evictionCache.Add("key3", "value3"))
	assert.Equal(t, "value1", evictionCache.Get("key1"))
	assert.Equal(t, "value2", evictionCache.Get("key2"))
	assert.Equal(t, 0, evictionCache.Add("key4", "value4"))
	assert.ElementsMatch(t, []string{"key1", "key2", "key4"}, evictionCache.Keys())
	assert.Equal(t, 0, evictionCache.Add("key3", "value3"))
	assert.ElementsMatch(t, []string{"key1", "key2", "key3"}, evictionCache.Keys())
	assert.Equal(t, 0, evictionCache.Add("key4", "value4"))
	assert.ElementsMatch(t, []string{"key1", "key2", "key4"}, evictionCache.Keys())
}
