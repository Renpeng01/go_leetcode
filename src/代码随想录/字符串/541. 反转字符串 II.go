package main

func reverseStr(s string, k int) string {

	if k == 1 {
		return s
	}

	bs := []byte(s)

	n := len(s) / (2 * k)

	for i := 0; i < n; i++ {
		left := i * (2 * k)
		right := i*(2*k) + k - 1
		for left < right {
			bs[left], bs[right] = bs[right], bs[left]
			left++
			right--
		}
	}

	if len(s)%(2*k) > 0 {
		left := (n - 1) * (2 * k)
		right := (n-1)*(2*k) + k - 1
		if len(s)%(2*k) <= k {
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
