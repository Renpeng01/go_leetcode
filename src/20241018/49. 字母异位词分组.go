package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	sameWords := make(map[string][]string)
	for _, w := range strs {

		key := []byte(w)
		sort.SliceStable(key, func(i, j int) bool {
			return key[i] < key[j]
		})

		_, ok := sameWords[string(key)]
		if !ok {
			sameWords[string(key)] = make([]string, 0, 16)
		}
		sameWords[string(key)] = append(sameWords[string(key)], w)
	}

	res := make([][]string, 0, len(sameWords))
	for _, v := range sameWords {
		res = append(res, v)
	}
	return res
}

func main() {
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	for _, w := range words {
		key := []byte(w)
		sort.SliceStable(key, func(i, j int) bool {
			return key[i] < key[j]
		})
		fmt.Println(string(key), w)
	}
}
