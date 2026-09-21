func jump(nums []int) int {
	jumps := 0 
	maxReachable := 0 
	currentEnd := 0 

	for i :=0 ; i < len(nums) - 1; i++ {
		
		maxReachable = max(maxReachable,i+nums[i])
		
		if i == currentEnd{
			jumps++
			currentEnd = maxReachable
		}

	}
	return jumps 
    
}
