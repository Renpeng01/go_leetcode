package main

import "fmt"

func maxSubarraySumCircular(nums []int) int {
	newNums := make([]int, 0, len(nums)*2)
	dp := make([]int, 0, len(newNums)*2)
	for i := 0; i < 2; i++ {
		for _, v := range nums {
			newNums = append(newNums, v)
			dp = append(dp, v)
		}
	}

	// fmt.Println(newNums)

	windowSize := len(nums)
	l := 0
	r := 0
	val := newNums[0]

	for i := 1; i < len(newNums); i++ {
		r = i
		dp[i] = max(dp[i], dp[i-1]+newNums[i])
		// fmt.Println("111dp[i]:", dp[i])
		if r-l+1 <= windowSize {
			if dp[i] <= 0 {
				l = i + 1
				r = i + 1
			}
		} else {
			dp[i] = dp[i] - newNums[l]
			l++
			for m := l; m < r; m++ {
				if newNums[m] < 0 {
					dp[i] = dp[i] - newNums[m]
					l++
				} else {
					break
				}
			}

		}
		// fmt.Println("222dp[i]:", dp[i])
		val = max(dp[i], val)

	}
	return val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{-3, -2, -3}
	res := maxSubarraySumCircular(nums)
	// fmt.Printlln(res)
	fmt.Println(res)
}
