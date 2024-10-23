package main

func longestConsecutive(nums []int) int {
	set := make(map[int]bool, len(nums))
	for _, v := range nums {
		set[v] = true
	}

	max := 0
	for v := range set {
		if !set[v-1] {
			cur := v
			cnt := 1

			for set[cur+1] {
				cur++
				cnt++
			}

			if cnt > max {
				max = cnt
			}
		}
	}
	return max
}

// func main() {
// 	// data := []int{0, 3, 7, 2, 5, 8, 4, 6, 1}
// 	data := []int{100, 4, 200, 1, 3, 2}
// 	res := longestConsecutive(data)
// 	fmt.Println(res)
// }
