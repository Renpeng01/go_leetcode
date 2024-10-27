package pipline

import (
	"fmt"
	"math"
	"sync"
)

func echo(nums []int) <-chan int {

	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func odd(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			if n%2 != 0 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		sum := 0
		for n := range in {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

// func main(){
// 	var nums = []int{1,2,3,4}
// 	for n := range sum(sq(odd(echo(nums)))){
// 		fmt.Println(n)
// 	}
// }

// Fan in/out
// 一对多和多对一的pipline

func is_prime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 2 {
			return false
		}
	}
	return value > 1
}

func prime(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if is_prime(n) {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)

	for i := range a {
		a[i] = min + 1
	}
	return a
}

func merger(cs []<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
		}(c)
		wg.Done()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// 通过并发的方式对一个很长的数组中的质数进行求和运算，我们想先把数组分段求和，然后再把它们集中起来。
func main() {
	nums := makeRange(1, 1000000)

	in := echo(nums)

	var chans [5]<-chan int

	for i := range chans {
		chans[i] = sum(prime(in))
	}
	for n := range sum(merger(chans[:])) {
		fmt.Println(n)
	}
}
