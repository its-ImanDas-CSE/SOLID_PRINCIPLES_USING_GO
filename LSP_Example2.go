/* Vehicle class: This class might have a method called move(), which makes the vehicle move.
   Car class: This is a subclass of Vehicle, and it has its own version of the move() method, which might make the car move differently from a general vehicle.
   **According to LSP, if you use a Vehicle in your program and later replace it with a Car, everything should still behave as expected — the program should still work properly. */

package main

import "fmt"

// Define a Vehicle interface
type Vehicle interface { // An interface is a way to define a set of methods that any type (like a struct) must implement if it wants to be considered as that interface.
	move()
}

// Define a Vehicle type with a move method
type VehicleImpl struct{} // Defines an empty struct VehicleImpl. It will implement the Vehicle interface.

func (v VehicleImpl) move() { // Implements the move method for the VehicleImpl struct. This satisfies the Vehicle interface.
	fmt.Println("The vehicle is moving.") // Prints "The vehicle is moving." when the move method is called.
}

// Define a Car type that embeds Vehicle
type Car struct { // Defines a struct Car that embeds the VehicleImpl struct. This means Car inherits the fields and methods of VehicleImpl
	VehicleImpl
}

func (c Car) move() { // Defines a move method specifically for the Car struct, overriding the move method of VehicleImpl.
	fmt.Println("The car is driving.") // Prints "The car is driving." when the Car's move method is called.
}

// Function to start the trip
func startTrip(v Vehicle) { // Declares a function startTrip that accepts any type implementing the Vehicle interface.
	v.move() // Calls the move method of the Vehicle interface
}

func lsp_ex() {
	// Create a Vehicle object
	vehicle := VehicleImpl{}

	// Create a Car object
	car := Car{}

	// Start trips with both objects
	startTrip(vehicle) // Output: The vehicle is moving.
	startTrip(car)     // Output: The car is driving.
}
