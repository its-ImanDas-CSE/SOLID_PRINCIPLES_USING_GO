// ------------------------------------Interface Segregation Principle (ISP)------------------------------------

// Break big interfaces into smaller, more specific ones.
// In Go, interfaces should be small and specific to what the type can do.

package main

import (
	"fmt"
)

// Bad Example, where the interface is large.
type Animal interface {
	Speak() string
	Swim() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Bark"
}

func (d Dog) Swim() string {
	return "Dog can swim"
}

// Good Example
// Separate interfaces for swimming and speaking.
type Speaker interface {
	Speak() string
}

type Swimmer interface {
	Swim() string
}

type Dog2 struct{}

func (d Dog2) Speak() string {
	return "Bark"
}

func (d Dog2) Swim() string {
	return "Dog can swim"
}
func isp2() {
	dog2 := Dog2{}
	var speaker Speaker
	var swimmer Swimmer
	speaker = dog2
	swimmer = dog2
	fmt.Println(speaker.Speak())
	fmt.Println(swimmer.Swim())
}
