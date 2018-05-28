package main

import "fmt"

func main() {
	ret := twoSum([]int{1, 2, 3, 4}, 6)
	fmt.Printf("%v", ret)
}

func twoSum(nums []int, target int) []int {
	var numMap = make(map[int]int, len(nums))

	for i, value := range nums {

		if j, ok := numMap[target-value]; ok {
			return []int{j, i}
		}
		numMap[value] = i
	}
	return nil
}
