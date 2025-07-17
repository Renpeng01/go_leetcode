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
		pathSet := make(map[int]struct{}, 16)
		if isCircleByDfs(edgeNext, v, pathSet) {
			return false
		}
	}
	return true
}

func isCircleByDfs(edgeNext map[int][]int, curCourse int, pathSet map[int]struct{}) bool {
	if _, ok := pathSet[curCourse]; ok {
		return true
	}
	pathSet[curCourse] = struct{}{}

	for _, v := range edgeNext[curCourse] {
		if isCircleByDfs(edgeNext, v, pathSet) {
			return true
		}
	}

	return false
}

func main() {
	numCourses := 2
	prerequisites := [][]int{
		// {0, 1},
		// {1, 0},
		// {},
	}

	res := canFinish(numCourses, prerequisites)
	fmt.Println("res: ", res)

}
