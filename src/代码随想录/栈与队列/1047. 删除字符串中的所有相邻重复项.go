package main

func removeDuplicates(s string) string {
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
			continue
		}

		if stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, s[i])
	}
	return string(stack)
}
