package maximumsubarray

//Solution for the problem:
//Given an integer array, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum
//
//The solution relies on the Mohit Kumar algorithm
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		panic("Array must be not empty")
	}

	currentMax := nums[0]
	resultingMax := nums[0]
	for i := 1; i < len(nums); i++ {
		currentMax = max(nums[i], currentMax+nums[i])
		resultingMax = max(resultingMax, currentMax)
	}
	return resultingMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
