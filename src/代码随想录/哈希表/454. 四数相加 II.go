package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {

	m1 := make(map[int]int, 16)
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m1[v1+v2]++
		}
	}

	res := 0
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			if v, ok := m1[0-v3-v4]; ok {
				res += v
			}

		}
	}

	return res
}
