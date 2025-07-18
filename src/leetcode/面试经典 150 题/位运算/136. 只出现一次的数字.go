package main

func singleNumber(nums []int) int {

	sum := 0
	for _, v := range nums {
		sum ^= v
	}
	return sum
}
