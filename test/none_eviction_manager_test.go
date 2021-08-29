package evictionmanager

import (
	"testing"

	"github.com/in-memory-cache/cache"
	"github.com/in-memory-cache/evictionmanager"
	"github.com/stretchr/testify/assert"
)

func TestNoneNormalCase(t *testing.T) {
	noneEvictionManager := evictionmanager.NewNoneEvictionManager()
	evictionCache := cache.NewInMemoryCache(3, noneEvictionManager)
	assert.Equal(t, 0, evictionCache.Add("key1", "value1"))
	assert.Equal(t, 0, evictionCache.Add("key2", "value2"))
	assert.Equal(t, 0, evictionCache.Add("key3", "value3"))
	assert.Equal(t, 1, evictionCache.Add("key2", "value2.1"))
	assert.Equal(t, "value3", evictionCache.Get("key3"))
	assert.Equal(t, "value1", evictionCache.Get("key1"))
	assert.Equal(t, "value3", evictionCache.Get("key3"))
	assert.ElementsMatch(t, []string{"key1", "key2", "key3"}, evictionCache.Keys())
	assert.Panics(t, func() { evictionCache.Add("key4", "value4") })
	assert.ElementsMatch(t, []string{"key1", "key2", "key3"}, evictionCache.Keys())
	assert.Equal(t, 3, evictionCache.Clear())
	assert.ElementsMatch(t, []string{}, evictionCache.Keys())
}

func TestNoneLimit1(t *testing.T) {
	noneEvictionManager := evictionmanager.NewNoneEvictionManager()
	evictionCache := cache.NewInMemoryCache(1, noneEvictionManager)
	assert.Equal(t, 0, evictionCache.Add("key1", "value1"))
	assert.Panics(t, func() { evictionCache.Add("key2", "value2") })
	assert.Panics(t, func() { evictionCache.Add("key3", "value3") })
	assert.Panics(t, func() { evictionCache.Add("key2", "value2.1") })
	assert.Equal(t, "", evictionCache.Get("key3"))
	assert.Equal(t, "value1", evictionCache.Get("key1"))
	assert.Equal(t, "", evictionCache.Get("key3"))
	assert.ElementsMatch(t, []string{"key1"}, evictionCache.Keys())
	assert.Panics(t, func() { evictionCache.Add("key4", "value4") })
	assert.Equal(t, 1, evictionCache.Add("key1", "value1.1"))
	assert.ElementsMatch(t, []string{"key1"}, evictionCache.Keys())
	assert.Equal(t, 1, evictionCache.Clear())
	assert.ElementsMatch(t, []string{}, evictionCache.Keys())
}

func TestNoneClearNotFullKeys(t *testing.T) {
	noneEvictionManager := evictionmanager.NewNoneEvictionManager()
	evictionCache := cache.NewInMemoryCache(3, noneEvictionManager)
	assert.Equal(t, 0, evictionCache.Add("key1", "value1"))
	assert.Equal(t, 0, evictionCache.Add("key2", "value2"))
	assert.Equal(t, 2, evictionCache.Clear())
	assert.Equal(t, 0, evictionCache.Add("key1", "value1"))
	assert.Equal(t, 0, evictionCache.Add("key2", "value2"))
	assert.Equal(t, 1, evictionCache.Add("key2", "value3"))
	assert.ElementsMatch(t, []string{"key1", "key2"}, evictionCache.Keys())
	assert.Equal(t, 2, evictionCache.Clear())
	assert.Equal(t, 0, evictionCache.Add("key1", "value1"))
	assert.ElementsMatch(t, []string{"key1"}, evictionCache.Keys())
	assert.Equal(t, 1, evictionCache.Clear())
	assert.ElementsMatch(t, []string{}, evictionCache.Keys())
}
