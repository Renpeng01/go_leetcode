package main

import (
	"strconv"
	"strings"
)

func calculate(s string) int {
	ops := make([]byte, 0, 16)
	nums := make([]int, 0, 16)

	for i := 0; i < len(s); i++ {
		if s[i] == '-' || s[i] == '+' {
			ops = append(ops, s[i])
		} else if s[i] == '(' || s[i] == ')' || strings.TrimSpace(string(s[i])) == "" {
			continue
		} else {
			n, _ := strconv.Atoi(string(s[i]))
			if len(ops) == 0 {
				nums = append(nums, n)
			} else {
				if len(nums) == 0 {

				} else {
					n1 := nums[len(nums)-1]
					op := ops[len(ops)-1]

					nums = nums[:len(nums)-1]
					ops = ops[:len(ops)-1]
					nums = append(nums, calculateItem(n1, n, string(op)))
				}
			}
		}
	}

}

func calculateItem(n1, n2 int, op string) int {
	res := 0
	switch op {
	case "-":
		res = n1 - n2
	case "+":
		res = n1 + n2
	}
	return res

}
