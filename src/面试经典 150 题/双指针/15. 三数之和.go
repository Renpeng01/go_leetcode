package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	resMap := make(map[[3]int]struct{}, 8)
	sort.Ints(nums)
	for i, v := range nums {
		// fmt.Println(111111)
		r := [3]int{}
		left := 0
		right := len(nums) - 1
		for left < right {
			if left == i {
				left++
				// fmt.Println(22222)
				continue
			}

			if right == i {
				right--
				// fmt.Println(22222)
				continue
			}

			if nums[left]+nums[right] == -v {
				// fmt.Println(33333)
				// 排序
				if nums[i] < nums[left] {
					r[0] = nums[i]
					r[1] = nums[left]
					r[2] = nums[right]
				} else if nums[i] >= nums[left] && nums[i] < nums[right] {
					r[0] = nums[left]
					r[1] = nums[i]
					r[2] = nums[right]
				} else {
					r[0] = nums[left]
					r[1] = nums[right]
					r[2] = nums[i]
				}
				resMap[r] = struct{}{}
				r = [3]int{}
				if (right - left) >= 2 {
					if nums[right] == nums[right-1] {
						right--
					} else if nums[left] == nums[left+1] {
						left++
					} else {
						left++
						right--

					}
					continue
				}
				break

			}

			if nums[left]+nums[right] > -v {
				right--
			} else {
				left++
			}

		}
	}
	res := make([][]int, 0, len(resMap))
	for k := range resMap {
		item := make([]int, 0, 3)
		for _, v := range k {
			item = append(item, v)
		}
		res = append(res, item)
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)

}
