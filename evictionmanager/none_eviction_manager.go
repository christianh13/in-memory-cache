package evictionmanager

import (
	"errors"
)

type NoneEvictionManager struct {
	hashMap map[string]bool
}

func NewNoneEvictionManager() *NoneEvictionManager {
	return &NoneEvictionManager{hashMap: make(map[string]bool)}
}

func (evictionManager NoneEvictionManager) Push(key string) int {
	if _, exists := evictionManager.hashMap[key]; exists {
		return 1
	} else {
		evictionManager.hashMap[key] = true
		return 0
	}
}

// Returns an arbitrary key
func (evictionManager NoneEvictionManager) Pop() string {
	var poppedKey string
	for key := range evictionManager.hashMap {
		poppedKey = key
		break
	}
	delete(evictionManager.hashMap, poppedKey)
	return poppedKey
}

func (evictionManager NoneEvictionManager) Clear() int {
	deletedKeys := len(evictionManager.hashMap)
	for k := range evictionManager.hashMap {
		delete(evictionManager.hashMap, k)
	}
	return deletedKeys
}

func (evictionManager NoneEvictionManager) Use(key string) {
}

func (evictionManager NoneEvictionManager) HandleOverLimit() string {
	panic(errors.New("key_limit_exceeded"))
}
