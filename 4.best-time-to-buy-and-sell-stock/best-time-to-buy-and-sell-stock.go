package main

func main() {
	prices := []int{2,1,2,1,0,1,2}
	println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}

	max := 0
	tmp := max

	for i := 1; i < l; i++ {
		diff := prices[i] - prices[i-1]
		println(diff)
		if tmp <= 0 {
			tmp = diff
		} else {
			tmp += diff
		}

		if tmp > max {
			max = tmp
		}
	}
	return max
}
