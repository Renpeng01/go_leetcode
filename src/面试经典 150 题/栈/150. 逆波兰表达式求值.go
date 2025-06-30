package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {

	stack := make([]int, 0, len(tokens))
	for i := 0; i < len(tokens); i++ {
		if !operator(tokens[i]) {
			num, _ := strconv.Atoi(tokens[i])
			stack = append(stack, num)
			continue
		}
		a := stack[len(stack)-2]
		b := stack[len(stack)-1]

		res := 0
		switch tokens[i] {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		default:
			res = a / b
		}
		stack = stack[:len(stack)-1]
		stack[len(stack)-1] = res
	}

	return stack[0]
}

func operator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}

func main() {
	s := '0'
	a := 'a'
	fmt.Println(s, a)
}
