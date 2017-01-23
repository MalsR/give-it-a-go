package main

import (
	"fmt"
)

/* Guess you could say a struct is a container of some arbitory types or 
 * a way of commonly grouping a bunch of types. We can use structs in go 
 * for encapsulation
*/
type Aircraft struct {
    make string
    model string
    age float32
    airline string
    engines int
    maxPax int
    maxLoad float32
}

func main() {
    //common way of creating a new aircraft is by contructor literals, allocate memory and assign the data to this struct 
    emiratesA380 := Aircraft { make: "Airbus", model: "A380", age: 3.5, airline: "EK", engines: 4, maxPax: 500}
    emiratesA380.isHeavy()
    baA320 := Aircraft { make: "Airbus", model: "A320", age: 10.5, airline: "BA", engines: 2, maxPax: 500}
    baA320.isHeavy()
        
}

/*We achieve encapsulation in terms of behaviour with structs by telling a method the type it can
* be called from, only a struct of Aircraft can call isHeavy
*/
func (a Aircraft) isHeavy() {
    if a.engines > 2 {
        fmt.Println("Heavies inbound!")
    } else {
        fmt.Println("Not a Heavy!")
    }
}