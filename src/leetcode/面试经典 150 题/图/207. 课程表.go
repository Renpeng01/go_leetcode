package main

// 判断是否为合法的有向无环图
func canFinish(numCourses int, prerequisites [][]int) bool {

	if len(prerequisites) == 0 {
		return true
	}
	edgsNext := make(map[int][]int, 16)
	edgePre := make(map[int][]int, 16)
	for _, v := range prerequisites {
		cur := v[0]
		pre := v[1]
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
		return false
	}

	for _, v := range firsts {
		exist := make(map[int]bool, 16)
		if dfs(edgsNext, v, exist) {
			return false
		}
	}
	return true
}

func dfs(edgs map[int][]int, node int, exist map[int]bool) bool {
	if exist[node] {
		return true
	}

	exist[node] = true
	for _, v := range edgs[node] {
		if exist[v] {
			return true
		}

		if dfs(edgs, v, exist) {
			return true
		}
	}
	return false
}
