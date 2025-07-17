package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {

	if len(prerequisites) == 0 {
		return true
	}
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

	// fmt.Printf("edgeNext: %+v\n", edgeNext)
	// fmt.Printf("edgePre: %+v\n", edgePre)
	// fmt.Printf("startCourses: %+v\n", startCourses)

	if len(startCourses) == 0 {
		return false
	}

	for _, v := range startCourses {
		pathSet := make(map[int]bool, 16)
		dealSet := make(map[int]bool, 16)
		if isCircleByDfs(edgeNext, -1, v, pathSet, dealSet) {
			return false
		}

		for k := range dealSet {
			delete(courses, k)
		}
	}
	fmt.Printf("len(courses): %+v\n", len(courses))
	return len(courses) == 0
}

func isCircleByDfs(edgeNext map[int][]int, pre, curCourse int, pathSet, dealSet map[int]bool) bool {
	if pathSet[curCourse] {
		return true
	}
	pathSet[curCourse] = true
	dealSet[curCourse] = true

	for _, v := range edgeNext[curCourse] {
		if isCircleByDfs(edgeNext, curCourse, v, pathSet, dealSet) {
			return true
		}
		pathSet[v] = false
	}

	return false
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
