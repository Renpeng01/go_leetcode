package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {

	copy := make([]string, len(strs))
	for k, v := range strs {

		bytes := []byte(v)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		copy[k] = string(bytes)
	}

	// fmt.Println(copy)

	res := make([][]string, 0, len(strs)/2+1)
	existed := make(map[string]bool, len(copy)/2+1)
	for i := 0; i < len(copy); i++ {
		if existed[copy[i]] {
			continue
		}

		c := make([]string, 0, 8)
		c = append(c, strs[i])
		for j := i + 1; j < len(copy); j++ {
			if copy[i] == copy[j] {
				c = append(c, strs[j])
			}
		}
		res = append(res, c)
		existed[copy[i]] = true
	}
	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	fmt.Println(res)

}
