package cmddiffunctions_tree

import (
	"container/heap"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    NodeTree // The value of the item; arbitrary.
	priority int      // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}
type NodeTree struct {
	content     string
	weight      int
	left, right *NodeTree
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value NodeTree, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func MakeCodeTree(freq map[string]int) (*NodeTree, error) {
	pq := make(PriorityQueue, len(freq))
	{
		i := 0
		for value, priority := range freq {
			pq[i] = &Item{
				value: NodeTree{
					content: value,
					weight:  priority,
					left:    nil,
					right:   nil,
				},
				priority: priority,
				index:    i,
			}
			i++
		}
	}
	heap.Init(&pq)
	for len(pq) > 1 {
		left := pq.Pop().(*Item)
		right := pq.Pop().(*Item)
		new_item := &Item{
			value: NodeTree{
				content: "",
				weight:  right.value.weight + left.value.weight,
				left:    &left.value,
				right:   &right.value,
			},
			priority: right.value.weight + left.value.weight,
		}
		heap.Push(&pq, new_item)
		pq.update(new_item, new_item.value, new_item.priority)
	}
	return &pq.Pop().(*Item).value, nil
}
func GethufCode(root *NodeTree, target string, path string) (*NodeTree, string) {
	if root == nil {
		return nil, path
	}

	if root.content == target {
		return root, path
	}

	leftNode, leftPath := GethufCode(root.left, target, path+"0")
	if leftNode != nil {
		return leftNode, leftPath
	}

	rightNode, rightPath := GethufCode(root.right, target, path+"1")
	if rightNode != nil {
		return rightNode, rightPath
	}

	return nil, path
}
