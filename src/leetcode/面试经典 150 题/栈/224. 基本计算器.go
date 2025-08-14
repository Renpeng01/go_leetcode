package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calculate(s string) int {
	ops := make([]string, 0, 16)
	nums := make([]int, 0, 16)
	exps := buildExpression(s)

	for i := 0; i < len(exps); i++ {
		if exps[i] == "-" || exps[i] == "+" {
			ops = append(ops, exps[i])
		} else if exps[i] == "(" || exps[i] == ")" {
			continue
		} else {
			n, _ := strconv.Atoi(string(exps[i]))
			if len(ops) == 0 {
				nums = append(nums, n)
			} else {
				if len(nums) == 0 {
					nums = append(nums, -n)
				} else {
					n1 := nums[len(nums)-1]
					op := ops[len(ops)-1]

					nums = nums[:len(nums)-1]
					ops = ops[:len(ops)-1]
					nums = append(nums, calculateItem(n1, n, op))
				}
			}
		}
	}

	if len(nums) == 1 {
		return nums[0]
	}

	for len(nums) > 0 {
		n1 := nums[len(nums)-2]
		n2 := nums[len(nums)-1]
		op := ops[len(ops)-1]

		nums = nums[:len(nums)-2]
		ops = ops[:len(ops)-1]
		nums = append(nums, calculateItem(n1, n2, string(op)))
	}
	return nums[0]

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

func buildExpression(s string) []string {
	if len(s) == 0 {
		return []string{}
	}

	res := make([]string, 0, 32)
	if s[0] == '-' {
		s = "0" + s
	}

	isNegative := false
	for i := 0; i < len(s); {
		if len(strings.TrimSpace(string(s[i]))) == 0 {
			i++
			continue
		}

		if !isNum(s[i]) {
			if s[i] == '-' {
				isNegative = true
				res = append(res, "+")
			} else {
				res = append(res, string(s[i]))
			}

			i++
			continue
		}
		j := i
		for ; j < len(s); j++ {
			if !isNum(s[j]) {
				break
			}
		}
		if isNegative {
			res = append(res, "-"+s[i:j])
			isNegative = false
		} else {
			res = append(res, s[i:j])
		}

		i = j
		// fmt.Println(i, j)
	}
	return res
}

func isNum(c byte) bool {
	return c == '0' || c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9'
}

func main() {
	// s := "(11+(4+5+2)-3)+(6+8)"
	// s := "(1+(4+5+2)-3)+(6+8)"
	s := "1-(     -2))"
	res := buildExpression(s)
	fmt.Println(res)

}
