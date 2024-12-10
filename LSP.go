//----------------------------- Liskov Substitution Principle (LSP)-------------------------------------------------

// You should be able to replace an object with its subclass and still have valid behavior.
// In Go, we don't have classical inheritance, but we can still follow this principle using interfaces and struct embedding.

/* Imagine you have a "general" object, like a "Vehicle," and a more specific object, like a "Car," which is a type of "Vehicle" (so Car is a subclass of Vehicle).
The Liskov Substitution Principle says that you should be able to replace a "Vehicle" object with a "Car" object, and the program should still work correctly — without breaking anything.*/

// More detailed Example...
/* Vehicle class: This class might have a method called move(), which makes the vehicle move.
   Car class: This is a subclass of Vehicle, and it has its own version of the move() method, which might make the car move differently from a general vehicle.
   **According to LSP, if you use a Vehicle in your program and later replace it with a Car, everything should still behave as expected — the program should still work properly. */

package main

import (
	"fmt"
)

type Bird interface {
	Fly() string
}

type Sparrow struct{}

func (s Sparrow) Fly() string {
	return "Flying"
}

type Penguin struct{}

func (p Penguin) Fly() string {
	return "Cannot Fly"
}

func printFly(b Bird) {
	fmt.Println(b.Fly())
}

func lsp() {
	sparrow := Sparrow{}
	penguin := Penguin{}
	printFly(sparrow) // Works fine
	printFly(penguin) // This is fine too, but might not be desired behavior
}

// The Penguin can substitute Bird, but it doesn’t follow the expected behavior of Fly as Penguin can’t fly. This is a violation of LSP.
