package main

func subarraySum1(nums []int, k int) int {
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

func subarraySum(nums []int, k int) int {
	cnt := 0
	pre := 0
	preSum := make(map[int]int)
	preSum[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		s, ok := preSum[pre-k]
		if ok {
			cnt += s
		}
		preSum[pre]++

	}
	return cnt
}
