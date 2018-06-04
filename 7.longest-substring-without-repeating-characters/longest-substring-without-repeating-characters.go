package main

func main() {
	ret := lengthOfLongestSubstring("pwwkew")
	println(ret)
}

func lengthOfLongestSubstring(s string) int {

	location := [256]int{}
	for i := range location {
		location[i] = -1
	}

	maxLen, left := 0, 0

	for i := 0; i < len(s); i++ {
		// 如果未出现过，location[s[i]]为 -1 一定小于left
		if location[s[i]] >= left {
			left = location[s[i]] + 1 // 出现过，就从上次值+1
		} else if i+1-left > maxLen {
			// fmt.Println(s[left:i+1])
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}

	return maxLen
}
