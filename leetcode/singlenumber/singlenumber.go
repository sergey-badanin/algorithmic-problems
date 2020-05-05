package singlenumber

//Solution for the problem:
//Given a non-empty array of integers, every element appears twice except for one. Find that single one
//
//The solution relies on the properties of XOR logical operation: a XOR a = 0, Commutativity, Associativity
func singleNumber(nums []int) int {
	if len(nums)%2 == 0 {
		panic("Array has even number of elements")
	}

	result := 0
	for _, v := range nums {
		result ^= v
	}
	return result
}
