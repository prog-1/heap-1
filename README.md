# Exercise 1

Assume a min heap is defined based on a slice as follows:

```golang
type heap []int
```

> NOTE: The following invariant is maintained in the min heap: `h[parent(i)] <= h[i]`

Implement `func (h *heap) push(x int)` function that pushes the element `x` into
the heap. The time complexity must be $O(\log n)$, where $n$ is the size of the
heap.

Implement `func (h *heap) pop() int` function that removes and returns the minimum element from the heap. The time complexity must be $O(\log n)$, where $n$ is the size of the heap.

## Example

Assume a heap is represented by the following slice: `[1 3 9 7 5]`.

* After
pushing `4`, the heap is represented by the following slice: `[1 3 4 7 5 9]`.
* After popping the minimum element `1` from the heap `[1 3 4 7 5 9]`, the heap is represented by the following slice: `[3 5 4 7 9]`.

# Exercise 2

Solve "[Ingus koeficients](https://lio.lv/arhivs/arhivs2/2019_3_d1_uzd.pdf)" problem using a heap instead of a binary search tree.