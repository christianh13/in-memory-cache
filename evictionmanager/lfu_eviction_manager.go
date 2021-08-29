package evictionmanager

import (
	"container/heap"

	"github.com/in-memory-cache/datastructures"
)

type LFUEvictionManager struct {
	hashMap       map[string]*datastructures.Item
	frequencyHeap *datastructures.FrequencyHeap
}

func NewLFUEvictionManager() *LFUEvictionManager {
	return &LFUEvictionManager{hashMap: make(map[string]*datastructures.Item), frequencyHeap: &datastructures.FrequencyHeap{}}
}

func (evictionManager LFUEvictionManager) Push(key string) int {
	if _, exists := evictionManager.hashMap[key]; exists {
		return 1
	} else {
		item := datastructures.Item{Key: key, Frequency: 0, Index: len(evictionManager.frequencyHeap.Items)}
		heap.Push(evictionManager.frequencyHeap, &item)
		evictionManager.hashMap[key] = &item
		return 0
	}
}

func (evictionManager LFUEvictionManager) Pop() string {
	key := heap.Pop(evictionManager.frequencyHeap).(*datastructures.Item).Key
	delete(evictionManager.hashMap, key)
	return key
}

func (evictionManager LFUEvictionManager) Clear() int {
	deletedKeys := len(evictionManager.hashMap)
	for k := range evictionManager.hashMap {
		delete(evictionManager.hashMap, k)
	}
	evictionManager.frequencyHeap.Empty()
	return deletedKeys
}

func (evictionManager LFUEvictionManager) Use(key string) {
	item := evictionManager.hashMap[key]
	item.Frequency = item.Frequency + 1
	heap.Fix(evictionManager.frequencyHeap, item.Index)
}

func (evictionManager LFUEvictionManager) HandleOverLimit() string {
	return evictionManager.Pop()
}
