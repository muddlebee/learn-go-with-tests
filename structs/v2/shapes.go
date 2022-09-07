package main

type Object struct {
	Height float64
	Width  float64
}

// Perimeter returns the perimeter of a rectangle.
func Perimeter(object Object) float64 {
	return 2 * (object.Height + object.Width)
}

// Area returns the area of a rectangle.
func Area(object Object) float64 {
	return object.Height * object.Width
}
