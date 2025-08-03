package main

func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < 32; i++ {
		cnt := 0
		for _, v := range nums {
			cnt += ((v >> i) & 1)
		}

		if cnt%3 != 0 {
			res |= (1 << i)
		}

	}
	return res
}
