package main
import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(math.pi)  // doesn't work since `pi`, which starts from small letter, is not exported
	fmt.Println(math.Pi)  // does work :)
}
