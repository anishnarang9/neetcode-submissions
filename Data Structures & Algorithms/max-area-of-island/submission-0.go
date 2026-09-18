func maxAreaOfIsland(grid [][]int) int {
    var dfs func(int,int)int 
	maxArea := 0  

    dfs = func(r,c int)int{

        if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]){
            return 0
        }
        if grid[r][c] != 1 {
            return 0
        }
		
        grid[r][c] = 0

        return 1 + dfs(r-1,c) + dfs(r+1,c) + dfs(r,c-1) + dfs(r,c+1)
    
    }
    for r := 0 ; r < len(grid) ; r ++ {
        for c := 0 ; c < len(grid[0]); c++{
            if grid[r][c] == 1{
                curArea := dfs(r,c)
                if curArea > maxArea{
					maxArea = curArea
				}
            }
        }
    }
    return maxArea 
    
}
