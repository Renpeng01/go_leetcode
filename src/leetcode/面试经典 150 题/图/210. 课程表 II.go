package main

var valid bool

var result []int

func findOrder(numCourses int, prerequisites [][]int) []int {
	edgeNext := make(map[int][]int, 16)
	valid = true
	result = make([]int, 0, 16)
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
	if !valid {
		return []int{}
	}

	for i := 0; i < len(result)/2; i++ { // TODO 为什么交换
		result[i], result[numCourses-i-1] = result[numCourses-i-1], result[i]
	}

	return result
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
	result = append(result, curCourse)
}
