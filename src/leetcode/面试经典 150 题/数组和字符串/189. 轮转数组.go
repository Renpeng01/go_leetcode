package main

func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 || k%n == 0 {
		return
	}

	k = k % n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)

}
func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
