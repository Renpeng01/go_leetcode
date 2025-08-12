package main

import "math"

// https://www.bilibili.com/video/BV1ss5xzAEwR/?spm_id_from=333.337.search-card.all.click&vd_source=70c464e99440c207e9933663bb2e5166
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m := len(nums1)
	n := len(nums2)

	left := 0
	right := m

	for left <= right {
		i := (left + right) / 2
		j := (m+n+1)/2 - i

		left1 := math.MinInt
		if i != 0 {
			left1 = nums1[i-1]
		}
		right1 := math.MaxInt
		if i != m {
			right1 = nums1[i]
		}

		left2 := math.MinInt
		if j != 0 {
			left2 = nums2[j-1]
		}
		right2 := math.MaxInt
		if j != n {
			right2 = nums2[j]
		}

		if left1 <= right2 && left2 <= right1 {
			if (m+n)%2 == 1 {
				return float64(max(left1, left2))
			}
			return (float64(max(left1, left2)) + float64(min(right1, right2))) / float64(2)
		} else if left1 > right2 {
			right = i - 1
		} else {
			left = i + 1
		}
	}

	return 0

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
