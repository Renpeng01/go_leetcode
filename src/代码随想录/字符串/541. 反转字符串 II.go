package main

func reverseStr(s string, k int) string {

	if k == 1 {
		return s
	}

	bs := []byte(s)
	for i := 0; i < len(s); i = i + k - 1 {
		if (i+1)%k == 0 && ((i+1)/k)%2 == 1 {
			left := i - (k - 1)
			right := i
			for left < right {
				bs[left], bs[right] = bs[right], bs[left]
				left++
				right--
			}
		}
	}
	return string(bs)
}
