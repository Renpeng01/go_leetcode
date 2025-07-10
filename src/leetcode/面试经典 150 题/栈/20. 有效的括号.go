package main

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
		} else {
			if len(stack) <= 0 {
				return false
			}
			if s[i] == '}' {
				if stack[len(stack)-1] != '{' {
					return false
				}
				stack = stack[0 : len(stack)-1]
			} else if s[i] == ']' {
				if stack[len(stack)-1] != '[' {
					return false
				}
				stack = stack[0 : len(stack)-1]
			} else if s[i] == ')' {
				if stack[len(stack)-1] != '(' {
					return false
				}
				stack = stack[0 : len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}
