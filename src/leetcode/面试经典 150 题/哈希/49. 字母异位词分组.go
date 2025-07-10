package main

import (
	"fmt"
	"sort"
)

// 通过 但是性能较差
func groupAnagrams1(strs []string) [][]string {

	copy := make([]string, len(strs))
	for k, v := range strs {

		bytes := []byte(v)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		copy[k] = string(bytes)
	}

	// 下面2个for循环是性能问题的关键
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

func groupAnagrams(strs []string) [][]string {

	set := make(map[string][]string, len(strs))
	for _, v := range strs {
		bytes := []byte(v)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})

		set[string(bytes)] = append(set[string(bytes)], v)
	}
	res := make([][]string, 0, 8)
	for _, v := range set {
		res = append(res, v)
	}

	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	fmt.Println(res)

}
