package main
import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := 1.0 + 0i
	for i := 0; i< 50; i++ {
		z = z - (cmplx.Pow(z, 3) - x) / (3.0 * cmplx.Pow(z, 2))
	}
	return z
}

func main()  {
	fmt.Println("cube root of 2 = ", Cbrt(2))
	fmt.Println("(cube root of 2)^3 = ", cmplx.Pow(Cbrt(2), 3))
}
