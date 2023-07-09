package ds

// max heap vs min heap

/*
parent: (n-1)/2
left child: 2n + 1
right child: 2n + 2
*/

type MaxHeap struct {
	arr []int
}

func (h *MaxHeap) swap(k1, k2 int) {
	h.arr[k1], h.arr[k2] = h.arr[k2], h.arr[k1]
}

func left(key int) int {
	return 2*key + 1
}

func right(key int) int {
	return 2*key + 2
}

func parent(key int) int {
	return (key - 1) / 2
}

// to maxheapify wrt it's predecessors
func (h *MaxHeap) maxHeapifyUp(pos int) {
	// if index is zero then done
	// check if it is <=> parent
	// if less or equal then done - this is redundant
	for h.arr[pos] > h.arr[parent(pos)] {
		// if greater then swap it with that one and
		h.swap(pos, parent(pos))
		pos = parent(pos)
	}
}

// to mexaheapify wrt it's children
func (h *MaxHeap) maxHeapifyDown(position int) {

	last := len(h.arr) - 1
	l, r := left(position), right(position)
	childToCompare := 0
	for l <= last {
		if l == last {
			childToCompare = l
		} else if h.arr[l] > h.arr[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		if h.arr[position] < h.arr[childToCompare] {
			h.swap(position, childToCompare)
			position = childToCompare
			l, r = left(position), right(position)
		} else {
			return
		}
	}
}

// insert into the Heap
func (h *MaxHeap) Insert(key int) {
	h.arr = append(h.arr, key)
	h.maxHeapifyUp(len(h.arr) - 1)
}

// Extract the top node in the heap
func (h *MaxHeap) Extract() int {
	out := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.maxHeapifyDown(0)
	// remove the last node
	return out
}
