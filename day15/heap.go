package day15

import (
	"container/heap"
)

type Item struct {
	node     *Node
	priority int
	index    int
}

func NewMinHeap(node *Node) *MinHeap {
	var minheap MinHeap
	minheap = append(minheap, &Item{node: node, priority: node.cost})
	heap.Init(&minheap)
	return &minheap
}

type MinHeap []*Item

func (mh MinHeap) Len() int { return len(mh) }

func (mh MinHeap) Less(i, j int) bool {
	return mh[i].priority < mh[j].priority
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
	mh[i].index = i
	mh[j].index = j
}

func (mh *MinHeap) Push(x interface{}) {
	n := len(*mh)
	item := x.(*Item)
	item.index = n
	*mh = append(*mh, item)
}

func (mh *MinHeap) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(mh, item.index)
}

func (mh *MinHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*mh = old[0 : n-1]
	return item
}

func (mh *MinHeap) PushNode(n *Node, priority int) {
	heap.Push(mh, &Item{
		node:     n,
		priority: priority,
	})
}

func (mh *MinHeap) UpsertNode(n *Node, priority int) {
	item := mh.Find(n)
	if item != nil {
		mh.update(item, priority)
	} else {
		mh.PushNode(n, priority)
	}
}

func (mh *MinHeap) PopNode() *Node {
	return heap.Pop(mh).(*Item).node
}

func (mh *MinHeap) Find(n *Node) *Item {
	for _, item := range *mh {
		if n == item.node {
			return item
		}
	}
	return nil
}
