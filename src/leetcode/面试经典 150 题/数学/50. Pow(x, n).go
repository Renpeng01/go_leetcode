package main

import "math"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	isNegative := n < 0
	n = int(math.Abs(float64(n)))

	res := float64(1)
	if isNegative {
		for i := 1; i <= n; i++ {
			res = res / x
		}
	} else {
		for i := 1; i <= n; i++ {
			res = res * x
		}
	}

	return res
}
