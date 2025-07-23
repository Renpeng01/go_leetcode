package main

import "fmt"

var path []int
var res [][]int

func combine(n int, k int) [][]int {
	path = make([]int, 0, 2)
	res = make([][]int, 0, 16)
	backtracking(n, k, 1)
	return res

}

func backtracking(n, k, startIndex int) {
	if len(path) == k {

		// 注意，这里需要将path复制一份，不分res中每个元素都是同一个地址
		tmp := make([]int, 0, len(path))
		for _, v := range path {
			tmp = append(tmp, v)
		}
		res = append(res, tmp)
		return
	}

	for i := startIndex; i <= n; i++ {
		path = append(path, i)
		backtracking(n, k, i+1)
		path = path[:len(path)-1]

	}
}

func main() {

	a := combine(4, 2)
	fmt.Println("a", a)
}
