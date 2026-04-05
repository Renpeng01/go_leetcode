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
	// m, _ := strconv.Atoi(arr[1])

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

		fmt.Printf("node1: %+v, node2: %+v\n", node1, node2)
		grid[node1][node2] = 1
		grid[node2][node1] = 1
	}

	fmt.Println("grid---------------------")
	for _, v := range grid {
		fmt.Println(v)
	}

	dfs(grid, 0, n)

	fmt.Println("res---------------------")
	for _, v := range resut {
		fmt.Println(v)
	}

}

func dfs(grid [][]int, x, target int) {
	if x == target {
		resut = append(resut, path)
		return
	}

	for _, v := range grid[x] {
		if v == 1 {
			path = append(path, v)
			dfs(grid, v, target)
			path = path[:len(path)-1]
		}
	}
}
