package main

func rangeBitwiseAnd(left int, right int) int {

	shift := 0
	for left < right {
		left, right = left>>1, right>>1
		shift++
	}

	return right << shift

}
