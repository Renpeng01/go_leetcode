package main

import "fmt"

func findLength(nums1 []int, nums2 []int) int {

	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}

	// dp[i][j]: 以i-1为结尾的nums1 和 以 j-1为结尾的nums2 的最长重复子数组长度
	dp := make([][]int, len(nums1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums2)+1)
	}

	max := 0

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}

	return max
}

func main() {
	nums1 := []int{1, 1, 0, 0, 1}
	nums2 := []int{1, 1, 0, 0, 0}

	res := findLength(nums1, nums2)
	fmt.Println(res)
}
