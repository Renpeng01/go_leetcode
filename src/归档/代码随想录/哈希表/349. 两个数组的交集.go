package main

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{}, len(nums1))
	for _, v := range nums1 {
		set[v] = struct{}{}
	}

	res := make([]int, 0, len(nums2))
	exist := make(map[int]struct{}, len(nums2))
	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			if _, ok := exist[v]; !ok {
				res = append(res, v)
				exist[v] = struct{}{}
			}
		}
	}
	return res
}
