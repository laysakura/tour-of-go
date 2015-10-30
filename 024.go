package main
import (
	"fmt"
	"math"
)

func Sqrt_10loops(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (math.Pow(z, 2) - x) / (2 * z)
	}
	return z
}

func Sqrt_err_converge(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (math.Pow(z, 2) - x) / (2 * z)
		if math.Pow(math.Pow(z, 2) - x, 2) < 1e-20 { break; }
	}
	return z
}

func main() {
	fmt.Println("Sqrt_10loops err: ", math.Abs(Sqrt_10loops(2.0) - math.Sqrt(2.0)))
	fmt.Println("Sqrt_err_converge err: ", math.Abs(Sqrt_err_converge(2.0) - math.Sqrt(2.0)))
}
