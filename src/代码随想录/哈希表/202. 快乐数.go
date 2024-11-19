package main

func isHappy(n int) bool {
	set := make(map[int]struct{})
	for {
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n = n / 10
		}
		if sum == 1 {
			return true
		}

		if _, ok := set[sum]; ok {
			return false
		}
		set[sum] = struct{}{}
		n = sum
	}
}
