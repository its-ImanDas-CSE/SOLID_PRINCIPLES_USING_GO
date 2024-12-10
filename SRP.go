//-------------------------- Single Responsibility Principle (SRP) -----------------------------------------

// A class or function should have one responsibility—it should do one thing and do it well.
// In Go, this means that a function or struct should focus on one task and not try to handle everything.
package main

import (
	"fmt"
)

// Bad Example
// This function both adds two numbers and logs the operation. It has two responsibilities.
func addAndLog(a, b int) int {
	result := a + b
	fmt.Println("Adding", a, "and", b)
	return result
}

// Good Example
// This function only adds two numbers (one responsibility).
func add(a, b int) int {
	return a + b
}

// This function only logs (one responsibility).
func logOperation(a, b int) {
	fmt.Println("Adding", a, "and", b)
}

func main() {
	//ocp()
	lsp()
	//lsp_ex()
	//isp()
	//isp2()
	var a int = 5
	var b int = 2
	logOperation(a, b)           //Good Coding
	fmt.Println(add(a, b))       // Good Coding
	fmt.Println(addAndLog(a, b)) // Bad Coding
}

//In the good example, we separated responsibilities: one function does the adding, and another does the logging.
