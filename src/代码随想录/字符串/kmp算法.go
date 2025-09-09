package main

import "fmt"

// https://www.bilibili.com/video/BV1PD4y1o7nd/?vd_source=70c464e99440c207e9933663bb2e5166
// https://www.bilibili.com/video/BV1M5411j7Xx?vd_source=70c464e99440c207e9933663bb2e5166&spm_id_from=333.788.videopod.sections

func getNext(s string) []int {
	next := make([]int, len(s))
	// j 前缀末尾位置
	// i 后缀末尾位置
	j := 0
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
			next[i] = j
		}
	}
	return next
}

func main() {
	next := getNext("aabaaf")
	fmt.Println(next)
}
