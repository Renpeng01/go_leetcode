package main

import "fmt"

var valid bool

func canFinish(numCourses int, prerequisites [][]int) bool {
	edgeNext := make(map[int][]int, 16)
	valid = true
	for _, q := range prerequisites {
		cur := q[0]
		pre := q[1]

		if edgeNext[pre] == nil {
			edgeNext[pre] = make([]int, 0, 16)
		}
		edgeNext[pre] = append(edgeNext[pre], cur)
	}

	for v := 0; v < numCourses; v++ {
		if !valid {
			return false
		}
		pathSet := make(map[int]bool, 16)
		isCircleByDfs(edgeNext, v, pathSet)

	}
	return valid
}

func isCircleByDfs(edgeNext map[int][]int, curCourse int, pathSet map[int]bool) {
	if pathSet[curCourse] {
		valid = false
		return
	}
	pathSet[curCourse] = true
	for _, v := range edgeNext[curCourse] {
		if !valid {
			return
		}
		isCircleByDfs(edgeNext, v, pathSet)
		pathSet[v] = false
	}
}

func main() {
	numCourses := 2
	prerequisites := [][]int{
		{1, 0},
	}

	res := canFinish(numCourses, prerequisites)
	fmt.Println("res: ", res)

}
