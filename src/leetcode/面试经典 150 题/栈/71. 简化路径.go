package main

import (
	"strings"
)

func simplifyPath(path string) string {
	pathEles := strings.Split(path, "/")
	stack := make([]string, 0, len(pathEles))
	for _, v := range pathEles {
		if v == "" {
			continue
		}

		if v == "." {
			continue
		}

		if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, v)
	}

	return "/" + strings.Join(stack, "/")
}

// func main() {
// 	s := "///"

// 	r := strings.Split(s, "/")

// 	fmt.Println(r)

// }
