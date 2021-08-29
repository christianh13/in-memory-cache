package main

import (
	"fmt"

	"github.com/in-memory-cache/cache"
	"github.com/in-memory-cache/evictionmanager"
)

// Sample test cases in problem statement
func main() {
	fmt.Println("None Eviction Manager")
	noneEvictionManager := evictionmanager.NewNoneEvictionManager()
	noneEvictionCache := cache.NewInMemoryCache(3, noneEvictionManager)
	fmt.Println(noneEvictionCache.Add("key1", "value1"))
	fmt.Println(noneEvictionCache.Add("key2", "value2"))
	fmt.Println(noneEvictionCache.Add("key3", "value3"))
	fmt.Println(noneEvictionCache.Add("key2", "value2.1"))
	fmt.Println(noneEvictionCache.Get("key3"))
	fmt.Println(noneEvictionCache.Get("key1"))
	fmt.Println(noneEvictionCache.Get("key3"))
	fmt.Println(noneEvictionCache.Keys())
	// Uncomment the line below to try panic
	// fmt.Println(noneEvictionCache.Add("key4", "value4"))
	fmt.Println(noneEvictionCache.Keys())
	fmt.Println(noneEvictionCache.Clear())
	fmt.Println(noneEvictionCache.Keys())

	fmt.Println("LRU Eviction Manager")
	lruEvictionManager := evictionmanager.NewLRUEvictionManager()
	lruEvictionCache := cache.NewInMemoryCache(3, lruEvictionManager)
	fmt.Println(lruEvictionCache.Add("key1", "value1"))
	fmt.Println(lruEvictionCache.Add("key2", "value2"))
	fmt.Println(lruEvictionCache.Add("key3", "value3"))
	fmt.Println(lruEvictionCache.Add("key2", "value2.1"))
	fmt.Println(lruEvictionCache.Get("key3"))
	fmt.Println(lruEvictionCache.Get("key1"))
	fmt.Println(lruEvictionCache.Get("key2"))
	fmt.Println(lruEvictionCache.Keys())
	fmt.Println(lruEvictionCache.Add("key4", "value4"))
	fmt.Println(lruEvictionCache.Keys())
	fmt.Println(lruEvictionCache.Clear())
	fmt.Println(lruEvictionCache.Keys())

	fmt.Println("LFU Eviction Manager")
	lfuEvictionManager := evictionmanager.NewLFUEvictionManager()
	lfuEvictionCache := cache.NewInMemoryCache(3, lfuEvictionManager)
	fmt.Println(lfuEvictionCache.Add("key1", "value1"))
	fmt.Println(lfuEvictionCache.Add("key2", "value2"))
	fmt.Println(lfuEvictionCache.Add("key3", "value3"))
	fmt.Println(lfuEvictionCache.Add("key2", "value2.1"))
	fmt.Println(lfuEvictionCache.Get("key3"))
	fmt.Println(lfuEvictionCache.Get("key1"))
	fmt.Println(lfuEvictionCache.Get("key2"))
	fmt.Println(lfuEvictionCache.Get("key3"))
	fmt.Println(lfuEvictionCache.Get("key1"))
	fmt.Println(lfuEvictionCache.Keys())
	fmt.Println(lfuEvictionCache.Add("key4", "value4"))
	fmt.Println(lfuEvictionCache.Keys())
	fmt.Println(lfuEvictionCache.Clear())
	fmt.Println(lfuEvictionCache.Keys())
}
