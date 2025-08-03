package main

// 超时
func rangeBitwiseAnd1(left int, right int) int {
	res := left
	for i := left + 1; i <= right; i++ {
		res &= i
	}
	return res
}
