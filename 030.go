package main
import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}  // called a slice, not a array or vector!
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}
