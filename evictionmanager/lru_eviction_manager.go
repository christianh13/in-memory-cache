package evictionmanager

import (
	"container/list"
)

type LRUEvictionManager struct {
	hashMap          map[string]*list.Element
	doublyLinkedList *list.List
}

func NewLRUEvictionManager() *LRUEvictionManager {
	return &LRUEvictionManager{hashMap: make(map[string]*list.Element), doublyLinkedList: list.New()}
}

func (evictionManager LRUEvictionManager) Push(key string) int {
	if _, exists := evictionManager.hashMap[key]; exists {
		return 1
	} else {
		pushedElement := evictionManager.doublyLinkedList.PushBack(key)
		evictionManager.hashMap[key] = pushedElement
		return 0
	}
}

func (evictionManager LRUEvictionManager) Pop() string {
	element := evictionManager.doublyLinkedList.Front()
	evictionManager.doublyLinkedList.Remove(element)
	return element.Value.(string)
}

func (evictionManager LRUEvictionManager) Clear() int {
	deletedKeys := len(evictionManager.hashMap)
	for k := range evictionManager.hashMap {
		delete(evictionManager.hashMap, k)
	}
	evictionManager.doublyLinkedList.Init()
	return deletedKeys
}

func (evictionManager LRUEvictionManager) Use(key string) {
	element := evictionManager.hashMap[key]
	evictionManager.doublyLinkedList.Remove(element)
	evictionManager.doublyLinkedList.PushBack(key)
}

func (evictionManager LRUEvictionManager) HandleOverLimit() string {
	return evictionManager.Pop()
}
