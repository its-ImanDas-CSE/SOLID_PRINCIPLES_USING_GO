// This is the same Programm which i wrote in ISP.go
// The difference is there, I created interface variable inside main function, and stored struct variable in the interface variable and call the function.
// but here i created the function outside main, so in main i define struct variable and call the function, the outer two new function do the job of rest that we are doing inside main in previous ISP.go programm.
package main

import (
	"fmt"
)

// Bad Example, where the interface is large.
type Animall interface {
	Speakk() string
	Swimm() string
}

type Dogg struct{}

func (d Dogg) Speakk() string {
	return "Bark"
}

func (d Dogg) Swimm() string {
	return "Dog can swim"
}

// Good Example
// Separate interfaces for swimming and speaking.
type Speakerr interface {
	Speakk() string
}

type Swimmerr interface {
	Swimm() string
}

type Dog22 struct{}

func (d Dog22) Speakk() string {
	return "Bark"
}

func (d Dog22) Swimm() string {
	return "Dog can swim"
}

func speakerr(s Speakerr) {
	fmt.Println(s.Speakk())
}
func swimmerr(s Swimmerr) {
	fmt.Println(s.Swimm())
}
func isp() {
	dog22 := Dog22{}
	speakerr(dog22)
	swimmerr(dog22)
}
