package happynumber

//Solution for the problem:
//Write an algorithm to determine if a number is "happy"
//
//The solution relies on the Floyd’s Cycle-Finding Algorithm
func isHappy(n int) bool {
	slow := n
	fast := sumOfSquaredDigits(sumOfSquaredDigits(n))
	for fast != slow && fast > 1 {
		slow = sumOfSquaredDigits(slow)
		fast = sumOfSquaredDigits(sumOfSquaredDigits(fast))
	}
	return fast == 1
}

func sumOfSquaredDigits(num int) int {
	result := 0
	for num > 0 {
		digit := num % 10
		result += digit * digit
		num = num / 10
	}
	return result
}
