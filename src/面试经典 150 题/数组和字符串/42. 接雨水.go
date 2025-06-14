package main

import "fmt"

// 超时
func trap(height []int) int {
	base := 1
	total := 0
	for {
		b := -1
		heighterCnt := 0
		for i := 0; i < len(height); i++ {
			if height[i] >= base {
				if b < 0 {
					b = i
				} else {
					total += (i - b) - 1
					b = i
				}
				heighterCnt++
			}
			// fmt.Println("b: ", b)
		}
		// fmt.Println("currTotal: ", total, " b: ", b)
		if heighterCnt <= 1 {
			break
		}
		base++
	}
	return total
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap(height)
	fmt.Println(res)
}
