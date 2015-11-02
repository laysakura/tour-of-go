package main
import (
	"math"
	"fmt"
)

type Abser interface  {
	Abs() float64
}
func main()  {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = &f   // MyFloat implements Abs()
	fmt.Println(a.Abs())

	a = &v   // Vertex implements Abs()
	fmt.Println(a.Abs())
}

type MyFloat float64

func (f *MyFloat) Abs() float64 {
	return math.Sqrt(float64(*f) * float64(*f))
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
