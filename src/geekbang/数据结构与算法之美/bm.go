package main

import (
	"fmt"
	"math"
)

const SIZE = 256 // 表示一共有多少种字符

// a,b 表示主串和模式串 n,m表示主串和模式串的长度
func bm(a []byte, n int, b []byte, m int) int {
	bc := generateBC(b, m)
	suffix, prefix := generateGS(b, m)

	i := 0
	for i < n-m {
		j := 0

		for j := m - 1; j >= 0; j-- {
			if a[i+j] != b[j] {
				break
			}
		}
		if j < 0 {
			return i
		}

		x := j - bc[int(a[i+1])]
		y := 0

		if j < m-1 {
			y = moveByGS(j, m, suffix, prefix)
		}
		i = i + int(math.Max(float64(x), float64(y)))
	}
	return -1

}

func moveByGS(j, m int, suffix []int, prefix []bool) int {
	k := m - 1 - j
	if suffix[k] != -1 {
		return j - suffix[k] + 1
	}
	for r := j + 2; r <= m-1; r++ {
		if prefix[r] {
			return r
		}
	}
	return m
}

func generateBC(b []byte, m int) []int {
	bc := make([]int, SIZE)
	for i := 0; i < SIZE; i++ {
		bc[i] = -1
	}

	for i := 0; i < m; i++ {
		bc[int(b[i])] = i
	}
	return bc
}

func generateGS(b []byte, m int) ([]int, []bool) {
	suffix := make([]int, m)
	prefix := make([]bool, m)

	for i := 0; i < m; i++ {
		suffix[i] = -1
		prefix[i] = false
	}

	for i := 0; i <= m-1; i++ {
		j := i
		k := 0
		for j >= 0 && b[j] == b[m-1-k] {
			j--
			k++               // 公共子串长度
			suffix[k] = j + 1 // j+1表示公共后缀子串在b[0,i]中的起始下标
		}
		if j == -1 {
			prefix[k] = true
		}
	}

	return suffix, prefix
}

func main() {
	a := 'a'
	fmt.Println(int(a))
}
