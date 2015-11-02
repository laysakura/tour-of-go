package main
import "fmt"

func main()  {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	// c <- 3  // fatal error: all goroutines are asleep - deadlock!
	           // channel size exceeded...
	fmt.Println(<-c)
	fmt.Println(<-c)
}
