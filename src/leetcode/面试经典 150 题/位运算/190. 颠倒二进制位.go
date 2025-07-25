package main

import (
	"fmt"
	"math"
)

func reverseBits(n int) int {

	bits := make([]int, 0, 16)

	for n > 0 {
		a := n % 2
		n = n / 2
		bits = append(bits, a)
	}

	bitsLen := len(bits)
	for i := 0; i < 32-bitsLen; i++ {
		bits = append(bits, 0)
	}
	// fmt.Println(bits)

	res := 0
	for i := len(bits) - 1; i >= 0; i-- {
		if bits[len(bits)-1-i] != 0 {
			res += int(math.Pow(2, float64(i)))
		}
	}
	return res
}

func main() {
	res := reverseBits(43261596)
	fmt.Println(res)
}
