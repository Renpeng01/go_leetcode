package main

import "fmt"

var table map[int]int

func init() {
	table = make(map[int]int, 16)
}

func hammingWeight(n int) int {

	sum := 0
	if v, ok := table[n]; ok {
		return v
	}
	for n > 0 {
		a := n % 2
		n = n / 2
		if a == 1 {
			sum++
		}
	}
	return sum
}

func main() {
	res := hammingWeight(11)
	fmt.Println(res)
}
