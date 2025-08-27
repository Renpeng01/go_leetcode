package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	set := make(map[int]int, len(nums1)+len(nums2))

	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			set[v1+v2]++
		}
	}

	cnt := 0
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			if c, ok := set[0-v3-v4]; ok {
				cnt = cnt + c
			}
		}
	}
	return cnt
}
