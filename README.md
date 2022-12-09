# Exercise 1

Assume a min heap is defined based on a slice as follows:

```golang
type Heap []int

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Push(x int)
// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Pop() int
// Len returns the number of elements in the collection.
func (h *Heap) Len() int { return len(*h) }
```

Implement `Heap.Push` and `Heap.Pop` functions.

> NOTE: The following invariant is maintained in the min heap: `h[parent(i)] <= h[i]`

## Example

Assume a heap is represented by the following slice: `[1 3 9 7 5]`.

* After
pushing `4`, the heap is represented by the following slice: `[1 3 4 7 5 9]`.
* After popping the minimum element `1` from the heap `[1 3 4 7 5 9]`, the heap is represented by the following slice: `[3 5 4 7 9]`.

# Exercise 2

Solve "[Ingus koeficients](https://lio.lv/arhivs/arhivs2/2019_3_d1_uzd.pdf)" problem using a heap instead of a binary search tree.
