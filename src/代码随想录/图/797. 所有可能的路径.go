package main

var res [][]int
var existed map[int]struct{}
var path []int
var N int

func allPathsSourceTarget(graph [][]int) [][]int {
	res = make([][]int, 0, 16)
	existed = make(map[int]struct{}, 16)
	path = make([]int, 0, 16)

	// n为目的地
	N = -1
	for k, v := range graph {
		if k > N {
			N = k
		}
		for _, s := range v {
			if s > N {
				N = s
			}
		}
	}

	dfs(0, graph)
	return res
}

func dfs(cur int, graph [][]int) {
	if cur == N {
		path = append(path, cur)
		res = append(res, path)
		path = make([]int, 0, 16)
		return

	}
	_, ok := existed[cur]
	if len(graph[cur]) == 0 || ok {
		return
	}

	existed[cur] = struct{}{}
	for _, p := range graph[cur] {
		path = append(path, p)
		dfs(p, graph)
		path = path[:len(path)-1]
	}
	delete(existed, cur)
}
