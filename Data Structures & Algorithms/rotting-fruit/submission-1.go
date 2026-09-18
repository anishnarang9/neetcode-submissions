func orangesRotting(grid [][]int) int {
	fresh := 0
	queue := [][]int{}
	minutes:= 0 

	for r := 0; r < len(grid); r++{
		for c := 0; c < len(grid[0]); c++{
			if grid[r][c] == 1{
				fresh++
			}
			if grid[r][c] == 2{
				queue = append(queue,[]int{r,c})
			}
		}
	}

	for len(queue) > 0 && fresh > 0 {
		levelSize := len(queue)

		for i := 0 ; i < levelSize; i++{
			r := queue[0][0]
			c := queue[0][1]
			queue = queue[1:]

			if r+1 < len(grid) && grid[r+1][c] == 1{
				grid[r+1][c] = 2
				fresh --
				queue = append(queue,[]int{r+1,c})
			}
			if r-1 >= 0  && grid[r-1][c] == 1{
				grid[r-1][c] = 2
				fresh --
				queue = append(queue,[]int{r-1,c})
			}
			if c + 1 < len(grid[0]) && grid[r][c+1] == 1{
				grid[r][c+1] = 2
				fresh --
				queue = append(queue,[]int{r,c+1})
			}
			if c-1 >= 0 && grid[r][c-1] == 1{
				grid[r][c-1] = 2
				fresh --
				queue = append(queue,[]int{r,c-1})
			}
			

		}
		minutes++

	}
	if fresh > 0 {
		return -1 
	}
	return minutes 
    
}
