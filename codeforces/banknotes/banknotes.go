package main

import "fmt"

const maxInt = int(^uint(0) >> 1)

/*
 * There are only three types of banknotes in Byterland: one franc, ten francs and one hundred francs.
 * What minimum number of banknotes do you need to get exactly n francs?
 */

func main() {
	var amnt int
	fmt.Scanf("%d\n", &amnt)
	coins := []int{1, 10, 100}
	fmt.Println(calcMininmalChange(coins, amnt))
}

/*
 * Classic coin change algorithm
 */
func calcMininmalChange(coins []int, amount int) int {
	minChanges := make([]int, amount+1)

	minChanges[0] = 0
	for i := 1; i <= amount; i++ {
		minChanges[i] = maxInt
	}

	for value := 1; value <= amount; value++ {
		for _, coinValue := range coins {
			if value >= coinValue {
				if minChanges[value-coinValue] < maxInt && minChanges[value-coinValue]+1 < minChanges[value] {
					minChanges[value] = minChanges[value-coinValue] + 1
				}
			}
		}
	}
	return minChanges[amount]
}
