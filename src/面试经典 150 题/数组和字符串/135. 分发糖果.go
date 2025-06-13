package main

import (
	"fmt"
)

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}

	res := make([]int, len(ratings))
	res[0] = 1

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			res[i] = res[i-1] + 1
		} else {
			res[i] = 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			res[i] = max(res[i], res[i+1]+1)
		}
	}

	total := 0
	for _, v := range res {
		total += v
	}
	return total

}
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	ratings := []int{1, 3, 2, 2, 1}

	res := candy(ratings)
	fmt.Println("res: ", res)

}
