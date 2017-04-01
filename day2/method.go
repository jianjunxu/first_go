package main

import (
	"math"
	"fmt"
)

func main() {
	r1 := Rectangle{3, 5}
	r2 := Rectangle{12, 9}
	c1 := Circle{12}
	c2 := Circle{2}

	fmt.Printf("r1 area: %f\r\n", r1.area())
	fmt.Printf("r2 area: %f\r\n", r2.area())
	fmt.Printf("c1 area: %f\r\n", c1.area())
	fmt.Printf("c2 area: %f\r\n", c2.area())
}

type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}