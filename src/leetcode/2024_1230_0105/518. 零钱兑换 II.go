package main

import "fmt"

func change(amount int, coins []int) int {

	dp := make([]int, amount+1)
	dp[0] = 1

	// 得到的是组合  (每次遍历一个物品，不会有2个物品以先后不同的顺序出现的问题)
	for i := 0; i < len(coins); i++ { // 物品
		for j := coins[i]; j <= amount; j++ { // 背包
			dp[j] += dp[j-coins[i]]
		}
	}

	// // 得到的是排列
	// for j := 0; j <= amount; j++ { // 背包
	// 	for i := 0; i < len(coins); i++ { // 物品

	// 	}
	// }
	return dp[amount]
}

func main() {
	fmt.Println(change(3, []int{2}))

}
