package main
import (
	"math"
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint(float64(e), " is negative")
}

func Sqrt(f float64) (float64, error)  {
	if f < 0 {
		return f, ErrNegativeSqrt(f)
	}
	return math.Sqrt(f), nil
}

func main()  {
	fmt.Println(Sqrt(2))

	_, err := Sqrt(-2)
	fmt.Println(err)
}
