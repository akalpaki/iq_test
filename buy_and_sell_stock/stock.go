package stock

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell
*/

// My algo: go through array. For each day, go through array again (from idx+1) and calculate profits. If profit > maxProfit (stored out of loop),
// replace it; else continue. Return maxProfit.

// too slow

func findProfitMyAlgo(prices []int) int {
	maxProfit := 0
	pLen := len(prices)
	for i := 0; i < pLen; i++ {
		dayPrice := prices[i]
		pricesCp := make([]int, pLen)
		copy(pricesCp, prices)
		for j := i + 1; j < pLen; j++ {
			profit := pricesCp[j] - dayPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

// Iteration 1: I loop once. I check if current value smaller than minimum; if not, I check if v - min > maxprofit.
// Very memory efficient, not so much for speed

func findProfitIteration1(prices []int) int {
	minPrice := 0
	maxProfit := 0
	for i, v := range prices {
		if i == 0 {
			minPrice = v
			continue
		}
		if v < minPrice {
			minPrice = v
			continue
		}
		profit := v - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

// Iteration 2: Fixing initial values and order of operations

func findProfitIteration2(prices []int) int {
	min, profit := prices[0], 0 // avoid 1 iteration by assigning first element as min right away
	pricesLen := len(prices)
	for i := 1; i < pricesLen; i++ {
		diff := prices[i] - min
		if prices[i] < min {
			min = prices[i]
		} else if diff > profit {
			profit = diff
		}
	}
	return profit
}
