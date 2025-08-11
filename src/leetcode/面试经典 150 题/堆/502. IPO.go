package main

import (
	"fmt"
	"sort"
)

type Project struct {
	Profit  int
	Capital int
}

// 超时
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {

	projects := make([]*Project, 0, len(profits))
	for i := 0; i < len(profits); i++ {
		projects = append(projects, &Project{
			Profit:  profits[i],
			Capital: capital[i],
		})
	}

	sort.SliceStable(projects, func(i, j int) bool {
		return projects[i].Profit > projects[j].Profit
	})

	visited := make(map[int]struct{}, len(projects))
	for k > 0 {
		for i, v := range projects {
			if _, ok := visited[i]; ok {
				continue
			}

			if w < v.Capital {
				continue
			}
			w = w + v.Profit
			visited[i] = struct{}{}
			break
		}
		k--
	}

	return w
}

func main() {
	profits := []int{1, 2, 3}
	capital := []int{0, 1, 1}

	k := 2
	w := 0

	res := findMaximizedCapital(k, w, profits, capital)
	fmt.Println(res)

}
