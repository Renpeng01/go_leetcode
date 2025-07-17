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

	pathSet := make(map[int]int, 16)
	for v := 0; v < numCourses && valid; v++ {
		if pathSet[v] == 0 {
			isCircleByDfs(edgeNext, v, pathSet)
		}
	}
	return valid
}

func isCircleByDfs(edgeNext map[int][]int, curCourse int, pathSet map[int]int) {
	pathSet[curCourse] = 1
	for _, v := range edgeNext[curCourse] {
		if pathSet[v] == 0 {

			if !valid {
				return
			}
			isCircleByDfs(edgeNext, v, pathSet)
		} else if pathSet[v] == 1 {
			valid = false
			return
		}

	}
	pathSet[curCourse] = 2 //被2标记过的公共路径不在重复走
}

func main() {
	numCourses := 2
	prerequisites := [][]int{
		{1, 0},
	}

	res := canFinish(numCourses, prerequisites)
	fmt.Println("res: ", res)

}
