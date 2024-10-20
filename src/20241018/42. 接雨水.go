package main

// 超时
func trap(height []int) int {
	maxH := 0
	for _, v := range height {
		if v > maxH {
			maxH = v
		}
	}

	total := 0
	for i := 1; i <= maxH; i++ {
		start := -1
		target := false
		for j, v := range height {
			if v >= i {
				if target {
					total += i - start - 1
					start = j
					target = false
				} else {
					start = j
				}
			} else {
				if start != -1 {
					target = true
				}
			}
		}
	}
	return total
}
