package main

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
	}

	startCourses := make([]int, 0, 16)
	for k := range courses {
		if _, ok := edgePre[k]; !ok {
			startCourses = append(startCourses, k)
		}
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
