package main

func intersection(nums1 []int, nums2 []int) []int {

	m := make(map[int]bool, len(nums1))
	for _, v1 := range nums1 {
		m[v1] = false
	}

	res := make([]int, 0, len(nums2))
	for _, v2 := range nums2 {
		if v, ok := m[v2]; ok && !v {
			res = append(res, v2)
			m[v2] = true
		}
	}
	return res

}
