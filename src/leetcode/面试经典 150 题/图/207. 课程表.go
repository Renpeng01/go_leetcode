package main

import "fmt"

// 判断是否为合法的有向无环图
func canFinish(numCourses int, prerequisites [][]int) bool {

	if len(prerequisites) == 0 {
		return true
	}
	edgsNext := make(map[int][]int, 16)
	edgePre := make(map[int][]int, 16)
	allCourse := make(map[int]struct{}, 16)
	for _, v := range prerequisites {
		cur := v[0]
		pre := v[1]
		allCourse[cur] = struct{}{}
		allCourse[cur] = struct{}{}
		if edgsNext[pre] == nil {
			edgsNext[pre] = make([]int, 0, 8)
		}
		edgsNext[pre] = append(edgsNext[pre], cur)

		if edgsNext[cur] == nil {
			edgsNext[cur] = make([]int, 0, 8)
		}

		if edgePre[cur] == nil {
			edgePre[cur] = make([]int, 0, 8)
		}
		edgePre[cur] = append(edgePre[cur], pre)

		if edgePre[pre] == nil {
			edgePre[pre] = make([]int, 0, 8)
		}
	}

	firsts := make([]int, 0, 16)
	for cur, pres := range edgePre {
		if len(pres) == 0 {
			firsts = append(firsts, cur)
		}
	}
	if len(firsts) == 0 {
		fmt.Println(111111)
		return false
	}

	exist := make(map[int]bool, 16)
	if hasCircle(edgsNext, firsts[0], exist) {
		fmt.Println(2222)
		return false
	}

	for k := range allCourse {
		if !exist[k] {
			fmt.Println(3333)
			return false
		}
	}
	return true
}

func hasCircle(edgs map[int][]int, node int, exist map[int]bool) bool {
	if exist[node] {
		return true
	}
	exist[node] = true
	for _, v := range edgs[node] {
		if exist[v] {
			return true
		}
		if hasCircle(edgs, v, exist) {
			return true
		}
	}
	return false
}

func main() {
	// prerequisites := [][]int{{0, 1}}
	prerequisites := [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}
	numCourses := 5

	res := canFinish(numCourses, prerequisites)
	fmt.Println("res: ", res)

}
