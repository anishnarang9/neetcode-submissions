func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0 
	currGas := 0 

	start := 0 
	for i := 0; i < len(gas); i++{
		totalGas += (gas[i] - cost[i])
		currGas += (gas[i] - cost[i])


		if currGas < 0{
		start = i + 1 
		currGas = 0
		}

		
		}
		if totalGas < 0 {
		return -1
	}

	
    return start
}
