package main

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

	if len(s)%step > 0 {
		left := (n - 1) * step
		right := (n-1)*step + k - 1
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
