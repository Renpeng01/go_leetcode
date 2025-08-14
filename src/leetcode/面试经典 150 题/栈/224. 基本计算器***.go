package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calculate1(s string) int {
	ops := make([]string, 0, 16)
	nums := make([]int, 0, 16)
	exps := buildExpression(s)

	fmt.Println("exps:", exps)

	for i := 0; i < len(exps); i++ {
		if exps[i] == "-" || exps[i] == "+" {
			ops = append(ops, exps[i])
		} else if exps[i] == "(" || exps[i] == ")" {
			continue
		} else {
			n, _ := strconv.Atoi(string(exps[i]))
			if len(ops) == 0 || (i > 0 && exps[i-1] == "(") {
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
					// fmt.Println("nums:", nums)
					// fmt.Println("ops:", ops)
				}
			}
		}
	}

	// fmt.Println("****", nums, len(nums))

	if len(nums) == 1 {
		return nums[0]
	}

	for len(nums) >= 2 {

		// fmt.Println("&&&&&&", nums, len(nums))
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

	// fmt.Printf("n1: %+v, op: %+v, n2: %+v\n", n1, op, n2)
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

	newS := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		newS += string(s[i])
	}

	new2S := ""
	for i := 0; i < len(newS); i++ {

		if newS[i] == '-' {

			if i != 0 {
				if newS[i-1] == '(' {
					new2S += "0"
				} else {
					new2S += "+0"
				}
			} else {
				new2S += "0"
			}

		}
		new2S += string(newS[i])
	}

	s = new2S
	fmt.Println("s:", s)

	isNegative := false
	for i := 0; i < len(s); {

		if !isNum(s[i]) {
			if s[i] == '-' && isNum(s[i+1]) {
				isNegative = true

				if i != 0 {
					res = append(res, "+")
				}

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

func main() {
	// s := "(11+(4+5+2)-3)+(6+8)"
	// s := "(1+(4+5+2)-3)+(6+8)"
	// s := "1-(     -2)"
	s := "-2+ 1"
	res := calculate(s)
	fmt.Println(res)

}

// https://www.bilibili.com/video/BV1HyGgzjEf3/?spm_id_from=333.337.search-card.all.click&vd_source=70c464e99440c207e9933663bb2e5166
func calculate(s string) int {

	s = strings.Replace(s, " ", "", -1)

	nums := make([]int, 0, 16)
	ops := make([]byte, 0, 16)

	cal := func() {
		r := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		l := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]

		val := 0
		switch op {
		case '+':
			val = l + r
		case '-':
			val = l - r
		}
		nums = append(nums, val)
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			ops = append(ops, s[i])
		} else if s[i] == ')' {
			for ops[len(ops)-1] != '(' {
				cal()
			}
			ops = ops[:len(ops)-1]
		} else if isNum(s[i]) {

			v := int(s[i] - '0')
			for ; i+1 < len(s) && isNum(s[i+1]); i++ {
				v = int(s[i+1]-'0') + v*10
			}
			nums = append(nums, v)
		} else {
			if s[i] == '-' {
				if i == 0 || s[i-1] == '(' {
					nums = append(nums, 0)
				}
			}

			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				cal()
			}
			ops = append(ops, s[i])
		}
	}

	for len(ops) > 0 {
		cal()
	}
	return nums[0]

}

func isNum(c byte) bool {
	return c == '0' || c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9'
}
