package pipline

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

func sq(in <-chan int) <-chan int {]
	out := make(chan int)
	go func ()  {
		for n := range in{
			out <- n * n
		}
		close(out)
	}
	return out
}

func odd(in <-chan int) <-chan int{
	out := make(chan int)

	go func ()  {
		for n := range in{
			if n%2 != 0{
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func sum(int <-chan int) <-chan int{
	out := make(chan int)
	go func ()  {
		sum := 0
		for n := range in{
			sum += n
		}
		out <- sum
		close(out)
	}()
	return  out
}

func main(){
	var nums = []int{1,2,3,4}
	for n := range sum(sq(odd(echo(nums)))){
		fmt.Println(n)
	}
}