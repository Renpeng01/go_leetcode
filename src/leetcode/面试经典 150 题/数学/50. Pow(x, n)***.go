package main

import "math"

// https://www.bilibili.com/video/BV19c411H7Pq/?spm_id_from=333.337.search-card.all.click&vd_source=70c464e99440c207e9933663bb2e5166
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	isNegative := n < 0
	n = int(math.Abs(float64(n)))

	res := float64(1)

	base := x
	for n > 0 {
		if n&1 == 1 {
			res = res * base
		}

		base = base * base
		n = n >> 1
	}

	if isNegative {
		return 1 / res
	}
	return res
}
