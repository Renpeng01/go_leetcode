package main

import (
	"fmt"
	"sort"
)

// hash表
func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	m := make(map[int]int, 16)
	for i, v := range nums {
		m[v] = i
	}
	res := make([][]int, 0, 8)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if index, ok := m[0-nums[i]-nums[j]]; ok && index > j {
				res = append(res, []int{nums[i], nums[j], nums[index]})
			}
		}
	}
	return res
}

// 双指针
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0, 8)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			return res // 注意这里不能return nil
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
				right--
			}
		}
	}
	return res
}

func main() {
	nums := []int{-2, 0, 1, 1, 2}
	res := threeSum(nums)
	fmt.Println(res)
}
