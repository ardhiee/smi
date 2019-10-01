package smi

import "math"

// Rectangle has a dimension of a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle has only radius
type Circle struct {
	Radius float64
}

// Area returns the area of Circles
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter will calculate the rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Shape is an interface that have Area
type Shape interface {
	Area() float64
}
