package main

func longestConsecutive(nums []int) int {
	next := make(map[int]int, len(nums))
	pre := make(map[int]int, len(nums))

	for _, v := range nums {
		next[v] = v + 1
		pre[v] = v - 1
	}

	max := 0
	exist := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := exist[nums[i]]; ok {
			continue
		}

		currentMax := 1
		nextKey := nums[i]
		ok := false
		for {
			nextKey, ok = next[nextKey]
			if !ok {
				break
			}
			currentMax++
			nextKey++
			exist[nextKey] = struct{}{}
		}

		preKey := nums[i]
		for {
			preKey, ok = pre[preKey]
			if !ok {
				break
			}
			currentMax++
			preKey--
			exist[preKey] = struct{}{}
		}

		if currentMax > max {
			max = currentMax
		}
	}
	return max
}
