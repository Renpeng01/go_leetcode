package main

import "fmt"

// 这个方法会因为n过大而导致total溢出
func trailingZeroes1(n int) int {

	total := 1
	for i := n; i >= 1; i-- {
		total *= i
	}

	fmt.Println(total)

	zeroCnt := 0
	for total > 0 {
		if total%10 != 0 {
			break
		}
		zeroCnt++
		total = total / 10
	}
	return zeroCnt
}

// https://www.bilibili.com/video/BV1maTazKEAt/?spm_id_from=333.337.search-card.all.click&vd_source=70c464e99440c207e9933663bb2e5166
// https://www.bilibili.com/video/BV1dd4y1T71f/?spm_id_from=333.337.search-card.all.click&vd_source=70c464e99440c207e9933663bb2e5166
func trailingZeroes(n int) int {
	cnt := 0

	for n > 0 {
		cnt += n / 5
		n = n / 5
	}
	return cnt

}
