package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := currentTime()
	fmt.Println("Current Time is", currentTime)

	//In order to run command line do:  go run functions/main.go functions/function_return_examples.go from checked out dir
	demonstrateGoFunctions()

	//***Examples of function pass by value and reference
	var message string = "Hello Core!"
	//print message but attempt to update message and print again
	printMessage(message)
	attemptToUpdateMessage(message)
	printMessage(message)

	//'&' pass in the memory address as opposed to string
	updateMessage(&message)
	printMessage(message)

	//Example of calling method with variadic parameters
	variadicMethodExample("Boeing 777", "Boing 747", "Airbus A350", "DC1030")
}

func currentTime() string {
	return time.Now().String()
}

func printMessage(message string) {
	fmt.Println(message)
}

//Pass by reference, where we send a reference that points to the memory location of where the variable is stored
func updateMessage(message *string) {
	//*string basically tells the method to expect a memory location that points to a string variable, i.e. pointer

	//*message is required and not message, because we are telling the compiler to assign what is right to the
	//the memory location referenced by the pointer; where the variable message is stored
	*message = "Hello SpeedBird 777!"

	//remember just 'message' would print the memory address of where the variable is stored, updateMessage(&message)
	//&message sends the memory location. In order to print the value use *message that tells to get the value
	//referenced by the memory location, dereference the pointer
	fmt.Println("Prints memory localtion of message variable=",message)
	fmt.Println("Print value of mesage=", *message)
}

//Pass by value, send a copy of a variables value into a function hence protect it from being changed
func attemptToUpdateMessage(message string) {
	message = "Hello Mals"
}

/**
Example of a variadic parameter in go. The '...' signals the creation of a variadic parameter in go.
Apparently must always be last in the declaration of parameters.
When calling method with variadic params go sends the values of the string as a slice.
 */
func variadicMethodExample(messages ...string) {
	//Treat as any other slice in go https://blog.golang.org/go-slices-usage-and-internals
	for _, message := range messages {
		fmt.Println(message)
	}
}

