package main

var res [][]int
var path []int

func allPathsSourceTarget(graph [][]int) [][]int {
	res = make([][]int, 0, 16)

	path = make([]int, 0, 16)
	path = append(path, 0)
	dfs(0, graph)
	return res
}

func dfs(cur int, graph [][]int) {
	if cur == len(graph)-1 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return

	}

	for _, p := range graph[cur] {
		path = append(path, p)
		dfs(p, graph)
		path = path[:len(path)-1]
	}

}
