package main

import "fmt"

func isHappy(n int) bool {

	set := make(map[int]bool)

	for {
		sum := 0
		for n > 0 {
			t1 := n % 10
			n = n / 10
			sum += t1 * t1

		}

		n = sum
		fmt.Println(sum)
		if set[sum] {
			return false
		}

		set[sum] = true
		if sum == 1 {
			return true
		}
	}

	return false
}

func main() {
	res := isHappy(19)
	fmt.Println(res)

}

// 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
// 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
// 如果这个过程 结果为 1，那么这个数就是快乐数。x
// 如果 n 是 快乐数 就返回 true ；不是，则返回 false
