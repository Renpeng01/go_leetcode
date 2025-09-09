package main

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 {
			return false
		}

		switch s[i] {
		case ')':
			if stack[len(stack)-1] != '(' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		case ']':
			if stack[len(stack)-1] != '[' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		case '}':
			if stack[len(stack)-1] != '{' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}
