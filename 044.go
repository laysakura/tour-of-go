package main
import "fmt"

func _fib(n_ int) int {
	if n_ <= 1 { return n_ }
	return _fib(n_ - 1) + _fib(n_ - 2)
}

func fibonacci() func() int {
	n := 0
	return func() int {
		n += 1
		return _fib(n)
	}
}

func main()  {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
