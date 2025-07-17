package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	edgeNext := make(map[int][]int, 16)
	edgePre := make(map[int][]int, 16)

	courses := make(map[int]struct{}, 16)
	for _, q := range prerequisites {
		cur := q[0]
		pre := q[1]

		if edgeNext[pre] == nil {
			edgeNext[pre] = make([]int, 0, 16)
		}
		edgeNext[pre] = append(edgeNext[pre], cur)

		if edgePre[cur] == nil {
			edgePre[cur] = make([]int, 0, 16)
		}
		edgePre[cur] = append(edgePre[cur], pre)
		courses[cur] = struct{}{}
		courses[pre] = struct{}{}
	}

	startCourses := make([]int, 0, 16)
	for k := range courses {
		if _, ok := edgePre[k]; !ok {
			startCourses = append(startCourses, k)
		}
	}

	valid := false
	for _, v := range startCourses {
		pathSet := make(map[int]bool, 16)
		isCircleByDfs(edgeNext, v, pathSet, &valid)
		if !valid {
			return false
		}
	}
	return valid
}

func isCircleByDfs(edgeNext map[int][]int, curCourse int, pathSet map[int]bool, valid *bool) {
	if pathSet[curCourse] {
		*valid = false
		return
	}
	pathSet[curCourse] = true

	for _, v := range edgeNext[curCourse] {
		if !*valid {
			return
		}
		isCircleByDfs(edgeNext, v, pathSet, valid)
	}
}

func main() {
	numCourses := 2
	prerequisites := [][]int{
		// {0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4},
		{5, 5},
	}

	res := canFinish(numCourses, prerequisites)
	fmt.Println("res: ", res)

}
