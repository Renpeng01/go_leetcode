package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	cnt := 0
	minLen := math.MaxInt
	for _, s := range strs {
		if len(s) < minLen {
			minLen = len(s)
		}
	}

	for i := 0; i < minLen; i++ {
		k := strs[0][i]

		hasDiff := false
		for j := 1; j < len(strs); j++ {
			if k != strs[j][i] {
				hasDiff = true
				break
			}
		}

		if hasDiff {
			return strs[0][:cnt]
		}
		cnt++
	}
	return strs[0][:cnt+1]
}

func main() {
	strs := []string{"dog", "racecar", "car"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)
}
