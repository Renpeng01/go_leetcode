package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	sumNums := make([]int, 0, len(nums))
	sum := 0
	for _, v := range nums {
		if v == target {
			return 1
		}
		sum += v
		sumNums = append(sumNums, sum)
	}

	cnt := len(sumNums)
	for i, v := range sumNums {
		for j := i - 1; j >= 0; j-- {
			if v-sumNums[j] == target {
				if i-j < cnt {
					cnt = i - j
				}
			}
		}
	}

	if cnt == len(sumNums) {
		return 0
	}
	return cnt
}

func main() {
	data := []int{2, 3, 1, 2, 4, 3}
	target := 11
	res := minSubArrayLen(target, data)
	fmt.Println(res)
}
