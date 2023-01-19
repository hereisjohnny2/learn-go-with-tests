package geometry

import "math"

type Shape interface {
	CalculateArea() float64
}

type Rectangle struct {
	x float64
	y float64
}

type Circle struct {
	r float64
}

type Triangle struct {
	b float64
	h float64
}

func (r Rectangle) CalculatePerimeter() float64 {
	return 2 * (r.x + r.y)
}

func (c Circle) CalculateArea() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func (r Rectangle) CalculateArea() float64 {
	return r.x * r.y
}

func (t Triangle) CalculateArea() float64 {
	return t.b * t.h / 2.0
}
