package main

import "fmt"

func longestConsecutive(nums []int) int {
	set := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		set[v] = struct{}{}
	}

	i := 0
	max := 0
	for i < len(nums) {
		if _, ok := set[nums[i]-1]; ok {
			i++
			continue
		}
		j := i + 1
		step := 1
		for j < len(nums) {
			if _, ok := set[nums[i]+step]; ok {
				step++
			} else {
				if j-i > max {
					max = j - i
				}
				step = 0
			}
			j++
		}
		i = j + 1
	}
	return max
}

func main() {
	// data := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	data := []int{100, 4, 200, 1, 3, 2}
	res := longestConsecutive(data)
	fmt.Println(res)
}
