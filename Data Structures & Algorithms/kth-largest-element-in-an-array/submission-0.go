type MinHeap []int

func (h MinHeap) Len() int{ return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int){ h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func findKthLargest(nums []int, k int) int {

	h := &MinHeap{}
	heap.Init(h)

	for _,v := range nums{
		heap.Push(h,v)
		if len(*h) > k{
			heap.Pop(h)
		}
	}
	return (*h)[0]

}
