type MaxHeap []int

func (h MaxHeap) Len() int{ return len(h) }

func (h MaxHeap) Less(i, j int) bool {return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int){ h[i], h[j] = h[j], h[i] }

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
func leastInterval(tasks []byte, n int) int {
	h := &MaxHeap{}
	heap.Init(h)
	time := 0 
	counts := make([]int ,26)

	for _,v := range tasks{
		counts[v - 'A']++
	}

	for _,count := range counts{
		if count >0 {
			heap.Push(h,count)
		}
	}

	type Cooldown struct{
		count int 
		readyTime int 
	}

	queue := []Cooldown{}

	for len(*h) > 0 || len(queue) > 0 {
		time++

		if len(queue) > 0 && queue[0].readyTime <= time {
			item := queue[0]
			queue = queue[1:]
			heap.Push(h,item.count)
		}

		if len(*h)  > 0 {
			count := heap.Pop(h).(int)
			count--

			if count > 0 {
				queue = append(queue,Cooldown{
					count: count,
					readyTime: time + n + 1,
				})
			}
			

		}
		
	}
	return time 

	

	

}
