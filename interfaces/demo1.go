package interfaces

import (
	"fmt"
	"math"
)

type shape interface{
	area() float64
}

type rectangle struct{
	rectangle_width, rectangle_height float64
}

type circle struct{
	radius float64
}

func (r rectangle) area() float64{
	return r.rectangle_height * r.rectangle_width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func school(s shape)  {
	fmt.Println("Area : ", s.area())
}

func Demo1()  {
	r:= rectangle{
		rectangle_width: 10, rectangle_height: 6,
	}

	school(r)
}