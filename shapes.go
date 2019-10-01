package smi

// Rectangle has a dimension of a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter will calculate the rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area will calculate area of rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
