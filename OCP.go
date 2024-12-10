//----------------------------- Open Closed Principle (OCP)-----------------------------------------------------------

// This means you should be able to add new features to your code without changing the existing code.
//In go,this is typically done using interfaces and structs.

package main

import (
	"fmt"
)

// Bad Example
// Every time we need to add a new shape, we need to modify the area calculation function.
func calculateArea_BadWay(shape string, radius float64, side float64) float64 {
	if shape == "circle" {
		return 3.14 * radius * radius
	} else if shape == "square" {
		return side * side
	}
	return 0
}

// Good Example
// We define an interface, and then extend it by adding new shapes without changing the code.
type Shape interface { // here we declare a method in interface.
	Area() float64
}

type Circle struct { // here we create a custom dataType circle, having only radius.
	radius float64
}

func (c Circle) Area() float64 { // This Area function("that we declare in interface") return the area of circle.
	// here c is Argument variable, Circle is custom dataType, Name of the function is Area(){"That we declare in interface"},and having return type float.
	return 3.14 * c.radius * c.radius
}

type Square struct { // here we create a custom dataType Square, having only side.
	side float64
}

func (s Square) Area() float64 { // This Area function("that we declare in interface") return the area of square.
	// here s is Argument variable, Square is custom dataType, Name of the function is Area(){"That we declare in interface"},and having return type float.
	return s.side * s.side
}

func calculateArea_goodWay(s Shape) float64 { // here we calculating the Area in the good way.
	return s.Area()
}

func ocp() { // here if I put it as main, it is throwing error as it is already declared."Karishma HCL".
	// Create a Circle with radius 5
	circle := Circle{radius: 5}

	// Create a Square with side length 4
	square := Square{side: 4}

	// Calculate and print the area of the circle
	circleArea := calculateArea_goodWay(circle)
	fmt.Printf("Area of Circle: %.2f\n", circleArea)

	// Calculate and print the area of the square
	squareArea := calculateArea_goodWay(square)
	fmt.Printf("Area of Square: %.2f\n", squareArea)
}

// here in this code, we uses interface & struct so, that we can achieve OCP.
// In the good example, you can add new shapes (like a Rectangle) by just creating a new struct that implements the Shape interface. You don’t have to change the calculateArea function.
