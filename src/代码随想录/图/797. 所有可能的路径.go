package main

var res [][]int

func allPathsSourceTarget(graph [][]int) [][]int {
	res = make([][]int, 0, 16)
	path := make([]int, 0, 16)
	dfs(graph, 0, 0, &path)
	return res
}

func dfs(graph [][]int, i, j int, path *[]int) {
	if i > len(graph)-1 || j > len(graph[0])-1 || i < 0 || j < 0 {
		return
	}

	if graph[i][j] == 2 {
		return
	}

	if i == len(graph)-1 && j == len(graph[0]) {
		res = append(res, *path)
		return
	}

	graph[i][j] = 2
	directions := [][]int{
		{-1, 0}, // 上
		{1, 0},  // 下
		{0, -1}, // 左
		{0, 1},  // 右边
	}

	for _, v := range directions {
		*path = append(*path, []int{i, j}...)
		dfs(graph, i+v[0], j+v[1], path)
		*path = (*path)[:len(*path)-1]
	}
	graph[i][j] = 1
}

// func main() {
// 	a := []int{1, 2, 3}
// 	test(a)
// 	fmt.Println(a)
// }

// func test(a []int) {
// 	a = append(a, 4)
// }
