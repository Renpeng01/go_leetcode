package main

func mySqrt(x int) int {

	l := 1
	r := x/2 + 1

	for l <= r {
		mid := l + (r-l)/2
		if mid*mid == x || mid*mid < x && (mid+1)*(mid+1) > x {
			return mid
		} else if mid*mid < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return 0
}
