package main

import "fmt"

func longestConsecutive(nums []int) int {
	set := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		set[v] = struct{}{}
	}

	fmt.Println("set: ", set)

	i := 0
	max := 0
	for i < len(nums) {
		fmt.Printf("i :%+v nums[i]: %+v\n", i, nums[i])
		if _, ok := set[nums[i]-1]; ok {
			i++
			fmt.Printf("existed in set nums[i]: %+v, nums[i]-1: %+v\n", nums[i], nums[i]-1)
			continue
		}
		j := i + 1
		step := 1
		for j < len(nums) {
			if _, ok := set[nums[i]+step]; ok {
				step++
				if j == len(nums)-1 {
					if j-i-1 > max {
						max = j - i - 1
					}
				}
			} else {
				fmt.Printf("max: %+v, (j-i-1): %+v \n", max, j-i-1)
				if j-i-1 > max {
					max = j - i - 1

				}
				step = 0
			}
			j++
		}
		i = j
	}
	return max
}

func main() {
	data := []int{0, 3, 7, 2, 5, 8, 4, 6, 1}
	// data := []int{100, 4, 200, 1, 3, 2}
	res := longestConsecutive(data)
	fmt.Println(res)
}
