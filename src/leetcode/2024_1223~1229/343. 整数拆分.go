package main

import "fmt"

func integerBreak(n int) int {

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 0
	}

	if n == 2 {
		return 1
	}

	dp := make([]int, n+1) // 对i进行拆分，得到最大的乘积为dp[i]

	dp[1] = 0
	dp[2] = 1

	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; /*这里也可以用 j<i, 此处的 j <= i/2 是个优化，因为乘积最大往往是几个比较接近的数相乘得到的  */ j++ {
			dp[i] = max(j*dp[i-j], max(dp[i], j*(i-j)))
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("res: ", integerBreak(10))
}
