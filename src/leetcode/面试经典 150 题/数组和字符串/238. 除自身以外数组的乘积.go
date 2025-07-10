package main

func productExceptSelf(nums []int) []int {

	prefixProduct := make([]int, len(nums))
	suffixPruduct := make([]int, len(nums))

	prefix := 1
	prefixProduct[0] = 1
	for i := 1; i < len(nums); i++ {
		prefixProduct[i] = prefix * nums[i-1]
		prefix = prefix * nums[i-1]
	}

	suffix := 1
	suffixPruduct[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		suffixPruduct[i] = suffix * nums[i+1]
		suffix = suffix * nums[i+1]
	}

	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = prefixProduct[i] * suffixPruduct[i]
	}
	return res
}
