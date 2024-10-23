package main

import (
	"sort"
	"strconv"
)

func threeSum1(nums []int) [][]int {
	h := make(map[int]int, len(nums))
	for i, v := range nums {
		h[v] = i
	}

	res := make([][]int, 0, 16)
	existed := make(map[string]bool)
	for index, num := range nums {
		target := -num
		for i, v := range nums {
			j, ok := h[target-v]
			if ok && index != j && j != i && index != i {
				arr := []int{-target, v, target - v}
				sort.Ints(arr)

				key := ""
				for _, v := range arr {
					s := strconv.Itoa(v)
					key += ":" + s
				}
				if !existed[key] {
					res = append(res, arr)
					existed[key] = true
				}
			}
		}
	}
	return res
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0, len(nums)/3)

	exist := make(map[[3]int]bool)
	for i := 0; i <= len(nums)-3; i++ {
		head := i + 1
		tail := len(nums) - 1
		for head < tail {
			if nums[i]+nums[head]+nums[tail] == 0 {
				if !exist[[3]int{nums[i], nums[head], nums[tail]}] {
					res = append(res, []int{nums[i], nums[head], nums[tail]})
					exist[[3]int{nums[i], nums[head], nums[tail]}] = true
				}
				tail--
				head++
			} else if nums[i]+nums[head]+nums[tail] > 0 {
				tail--
			} else {
				head++
			}
		}
	}

	return res
}

// func main() {
// 	data := []int{-1, 0, 1, 2, -1, -4}

// 	res := threeSum(data)
// 	fmt.Println(res)

// }
