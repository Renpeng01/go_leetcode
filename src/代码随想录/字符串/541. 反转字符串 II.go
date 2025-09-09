package main

import "fmt"

func reverseStr(s string, k int) string {

	if k == 1 {
		return s
	}

	bs := []byte(s)
	step := 2 * k
	n := len(s) / step
	for i := 0; i < n; i++ {
		left := i * step
		right := i*step + k - 1
		for left < right {
			bs[left], bs[right] = bs[right], bs[left]
			left++
			right--
		}
	}
	// fmt.Println(string(bs), n)

	if len(s)%step > 0 {
		left := (n) * step
		right := (n)*step + k - 1
		if len(s)%step <= k {
			right = len(s) - 1
		}
		for left < right {
			bs[left], bs[right] = bs[right], bs[left]
			left++
			right--
		}
	}

	return string(bs)
}

func main() {
	s := "abcdefg"
	k := 2
	fmt.Println(reverseStr(s, k))
}
