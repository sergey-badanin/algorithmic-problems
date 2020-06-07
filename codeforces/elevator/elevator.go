package main

import "fmt"

/* Fill elevator. Optimize elevators rides
 * Three people weighting x, y, and z kg want to take an elevator.
 * The elevator can lift people with a total weight of no more than m kg.
 * What is the minimum number of rides they need to take all of them?
 */

func main() {
	var x, y, z, m int
	fmt.Scanf("%d %d %d %d\n", &x, &y, &z, &m)
	runs := fillElevator([]int{x, y, z}, m)
	fmt.Println(runs)
}

func fillElevator(psngrs []int, limit int) int {
	arr := make([]int, len(psngrs))
	copy(arr, psngrs)
	steps := 0

	for len(arr) > 0 {
		steps++

		taken := make(map[int]bool)
		answ := make(map[int]bool)
		max := 0
		findMaxSum(arr, limit, taken, 0, &max, answ)

		var arr1 []int
		for i := range arr {
			if answ[i] == false {
				arr1 = append(arr[:i], arr[i+1:]...)
			}
		}
		arr = arr1
	}
	return steps
}

func findMaxSum(arr []int, limit int, taken map[int]bool, sum int, max *int, answ map[int]bool) {
	for i := 0; i < len(arr); i++ {
		if taken[i] == false {
			taken[i] = true
			if sum+arr[i] <= limit && sum+arr[i] > *max {
				*max = sum + arr[i]
				copyMap(taken, answ)
			}
			findMaxSum(arr, limit, taken, sum+arr[i], max, answ)
			taken[i] = false
		}
	}
}

func copyMap(src, tgt map[int]bool) {
	for k, v := range src {
		tgt[k] = v
	}
}
