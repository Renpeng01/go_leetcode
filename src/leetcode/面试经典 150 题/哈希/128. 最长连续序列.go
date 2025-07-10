package main

import "fmt"

func longestConsecutive1(nums []int) int {
	next := make(map[int]int, len(nums))
	pre := make(map[int]int, len(nums))

	// 优化其实不需要记录pre和next,用一个map即可  因为其next的val一定是key+1  pre 的val一定是key-1
	for _, v := range nums {
		next[v] = v + 1
		pre[v] = v - 1
	}

	fmt.Println(next, pre)

	max := 0
	exist := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := exist[nums[i]]; ok {
			continue
		}
		exist[nums[i]] = struct{}{}

		currentMax := 1
		nextKey, ok := next[nums[i]]
		for {
			nextKey, ok = next[nextKey]
			if !ok {
				break
			}
			currentMax++
			exist[nextKey] = struct{}{}
		}

		preKey, ok := pre[nums[i]]
		for {
			preKey, ok = pre[preKey]
			if !ok {
				break
			}
			currentMax++
			exist[preKey] = struct{}{}
		}

		if currentMax > max {
			max = currentMax
		}
	}
	return max
}
func longestConsecutive(nums []int) int {
	numsMap := make(map[int]bool, len(nums))

	for _, v := range nums {
		numsMap[v] = false
	}

	max := 0
	for _, v := range nums {
		if numsMap[v] {
			continue
		}

		numsMap[v] = true
		currentMax := 1
		nextkey := v + 1
		for {
			if _, next := numsMap[nextkey]; !next {
				break
			}
			numsMap[nextkey] = true
			nextkey++
			currentMax++
		}

		preKey := v - 1
		for {
			if _, pre := numsMap[preKey]; !pre {
				break
			}
			numsMap[preKey] = true
			preKey--
			currentMax++
		}

		if currentMax > max {
			max = currentMax
		}
	}
	return max
}
func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	res := longestConsecutive(nums)
	fmt.Println(res)
}
