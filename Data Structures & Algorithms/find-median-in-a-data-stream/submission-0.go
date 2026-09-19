
type MedianFinder struct {
	minh MinHeap
	maxh MaxHeap 
}

func Constructor() MedianFinder {
	minh := MinHeap{}
	heap.Init(&minh)
	maxh := MaxHeap{}
	heap.Init(&maxh)

	return MedianFinder {
		minh: minh,
		maxh: maxh,
	}
    
}


func (this *MedianFinder) AddNum(num int)  {
	if len(this.maxh) == 0{
		heap.Push(&this.maxh, num)
	}else{
		if num > this.maxh[0]{
			heap.Push(&this.minh, num)
		}else{
			heap.Push(&this.maxh, num)
		}
	}

	if len(this.maxh) > len(this.minh) + 1{
		temp := heap.Pop(&this.maxh)
		heap.Push(&this.minh, temp)

	}else if len(this.minh) > len(this.maxh) +1 {
		temp := heap.Pop(&this.minh)
		heap.Push(&this.maxh, temp)

	}
    
}


func (this *MedianFinder) FindMedian() float64 {
    if (len(this.minh) + len(this.maxh)) % 2 == 0 {
		return float64(this.minh[0]+this.maxh[0]) / 2.0
	}else{
		if len(this.maxh) > len(this.minh) {
    		return  float64(this.maxh[0])
		}else {
    		return  float64(this.minh[0])
		}
	}
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
