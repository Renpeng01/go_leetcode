package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var path []int
var resut [][]int

func main() {

	path = make([]int, 0, 16)
	resut = make([][]int, 0, 16)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arr := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(arr[0])

	grid := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		grid[i] = make([]int, n+1)
	}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		arr := strings.Split(line, " ")
		if len(arr) != 2 {
			fmt.Println()
			break
		}
		node1, _ := strconv.Atoi(arr[0])
		node2, _ := strconv.Atoi(arr[1])

		grid[node1][node2] = 1
	}

	path = append(path, 1)
	dfs(grid, 1, n)

	if len(resut) == 0 {
		fmt.Println(-1)
		return
	}

	for _, res := range resut {
		s := ""
		for i, v := range res {

			if i == 0 {
				s += strconv.Itoa(v)
				continue
			}
			s += (" " + strconv.Itoa(v))
		}
		fmt.Println(s)
	}

}

func dfs(grid [][]int, x, target int) {
	if x == target {
		new := make([]int, len(path))
		copy(new, path)
		resut = append(resut, new)
		return
	}

	for i := 1; i <= target; i++ {
		if grid[x][i] == 1 {
			path = append(path, i)
			dfs(grid, i, target)
			path = path[:len(path)-1]
		}
	}
}
