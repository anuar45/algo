package main

import "fmt"

func main() {

	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))

}

func maxProfit(prices []int) int {
	var totalprofit, maxprofit, profit int

	if len(prices) <= 1 {
		return 0
	}

	for i := 0; i < len(prices); i++ {

		maxprofit = 0

		for j := i + 1; j < len(prices); j++ {

			if prices[i] < prices[j] {
				//fmt.Println(i, j)
				profit = maxProfit(prices[j+1:]) + prices[j] - prices[i]

				if profit > maxprofit {
					maxprofit = profit
					//fmt.Println("maxprofit:", maxprofit)
				}
			}

		}

		if maxprofit > totalprofit {
			totalprofit = maxprofit
		}

	}
	return totalprofit
}
