package main

import (
	"fmt"
	"sort"
)

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	earnings := make([]int, len(profits))
	for i := 0; i < len(earnings); i++ {
		earnings[i] = profits[i] - capital[i]
	}
	fmt.Println("earnings", earnings)

	sort.Slice(capital, func(i, j int) bool {
		return earnings[i] > earnings[j]
	})
	fmt.Println("capital", capital)

	sort.Slice(earnings, func(i, j int) bool {
		return earnings[i] > earnings[j]
	})
	fmt.Println("newearnings", earnings)

	res := 0
	visited := make(map[int]struct{}, len(earnings))
	for k > 0 {
		for i, v := range capital {
			if _, ok := visited[i]; ok {
				continue
			}

			if w < v {
				continue
			}
			w = w + earnings[i]
			visited[i] = struct{}{}
			break
		}
		k--
	}

	return res
}

func main() {
	profits := []int{1, 2, 3}
	capital := []int{0, 1, 1}

	k := 2
	w := 0

	res := findMaximizedCapital(k, w, profits, capital)
	fmt.Println(res)

}
