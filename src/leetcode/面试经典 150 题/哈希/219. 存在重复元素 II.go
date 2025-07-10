package main

func containsNearbyDuplicate(nums []int, k int) bool {
	set := make(map[int]int, len(nums))
	for i, v := range nums {
		j, ok := set[v]
		if ok && abs(i, j) < k {
			return true
		}
		set[v] = i
	}
	return false
}

func abs(i, j int) int {
	if i > j {
		return i - j
	}

	return j - i
}
