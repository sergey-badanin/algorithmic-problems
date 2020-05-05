package buysellstock

//Solution for the problem:
//Say you have an array prices for which the ith element is the price of a given stock on day i.
//Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).
//Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).
//
// O(n) complexity solution
func maxProfit(prices []int) int {
	buyPrice := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-buyPrice > 0 {
			profit += prices[i] - buyPrice
		}
		buyPrice = prices[i]
	}
	return profit
}
