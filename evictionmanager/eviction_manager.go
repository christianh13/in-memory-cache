package evictionmanager

type EvictionManager interface {
	Push(key string) int
	Pop() string
	Clear() int
	Use(key string)
	HandleOverLimit() string
}
