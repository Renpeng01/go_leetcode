package main

func subarraySum(nums []int, k int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum == k {
				cnt++
			}
		}
	}
	return cnt
}
