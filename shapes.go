package main

import "math"

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
} //Rectangle has a method called Area that returns a float64 so it satisfies the Shape interface

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
} //Circle has a method called Area that returns a float64 so it satisfies the Shape interface

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return ((.5 * t.Base) * t.Height)
}
