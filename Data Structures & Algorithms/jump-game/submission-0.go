func canJump(nums []int) bool {
	maxReachable := 0 

	for i := 0 ; i < len(nums); i++{
		if i > maxReachable {
			return false 
		}	
		if i + nums[i] > maxReachable{
			maxReachable = i + nums[i]
		}
		
	}
	return true 
	
}
