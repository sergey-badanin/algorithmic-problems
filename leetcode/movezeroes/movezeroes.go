package movezeroes

//Solution for the problem:
//Given an array of integers, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements
//
// O(n) complexity solution
func moveZeroes(nums []int) {
	shift := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			shift++
		}
		if nums[i] != 0 && shift > 0 {
			nums[i-shift] = nums[i]
			nums[i] = 0
		}
	}
}
