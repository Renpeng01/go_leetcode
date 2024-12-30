package main

import "fmt"

//物品	重量	价值
// 0    1      15
// 1    3      20
// 3    4      30
// 背包最大承重量4

// dp[i][j] = max(dp[i-1][j],dp[i-1][j-weight[i]] + value[i])

// 二维数组
func bag2(bagV int, weight []int, value []int) int {
	dp := make([][]int, len(weight)) // dp[i][j]  0~i之间的物品，任取放进容量为j的背包里

	for i := 0; i < len(weight); i++ {
		dp[i] = make([]int, bagV+1)
	}

	for i := 1; i <= bagV; i++ {
		if i >= weight[0] {
			dp[0][i] = value[0]
		}
	}

	for i := 1; i < len(weight); i++ { // 遍历物品
		for j := 0; j <= bagV; j++ { // 遍历背包容量
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}

		}
	}

	fmt.Println(dp)

	return dp[len(weight)-1][bagV]
}

// 一维数组
func bag1(bagV int, weight []int, value []int) int {

	dp := make([]int, bagV+1) // dp[i] 容量为i的背包最大值为dp[i]
	// for i := 0; i <= bagV; i++ {
	// 	if i >= weight[0] {
	// 		dp[i] = value[0]
	// 	}
	// }
	// fmt.Println(dp)

	for i := 0; i < len(weight); i++ { // 遍历物品
		for j := bagV; j > weight[i]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
		fmt.Println(dp)
	}

	return dp[bagV]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}

	res := bag2(4, weight, value)
	fmt.Println("res: ", res)
}
