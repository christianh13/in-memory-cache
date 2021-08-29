package datastructures

type Item struct {
	Key       string
	Frequency int
	Index     int
}

type FrequencyHeap struct {
	Items []*Item
}

func (h FrequencyHeap) Len() int {
	return len(h.Items)
}

func (h FrequencyHeap) Less(i, j int) bool {
	return h.Items[i].Frequency < h.Items[j].Frequency
}

func (h FrequencyHeap) Swap(i, j int) {
	h.Items[i], h.Items[j] = h.Items[j], h.Items[i]
	h.Items[i].Index = i
	h.Items[j].Index = j
}

func (h *FrequencyHeap) Push(x interface{}) {
	n := len(h.Items)
	item := x.(*Item)
	item.Index = n
	h.Items = append(h.Items, item)
}

func (h *FrequencyHeap) Pop() interface{} {
	old := *h
	n := len(old.Items)
	x := old.Items[n-1]
	x.Index = -1
	h.Items = old.Items[0 : n-1]
	return x
}

func (h *FrequencyHeap) Empty() {
	h.Items = []*Item{}
}
