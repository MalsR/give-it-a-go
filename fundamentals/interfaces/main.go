package main

import (
	"fmt"
)

type airport struct {
	code        string
	city        string
	country     string
	dialingCode int
	terminals   int
}
//
//type Printer interface {
//	print() string
//}

func main() {
	heathrowAirport := airport{code: "LHR", city: "London", country: "UK", dialingCode: 44, terminals: 5}
	colomboAirport := airport{code: "CMB", city: "Katunayake", country: "SL", dialingCode: 94, terminals: 2}
	changiAirport := airport{code: "SGN", city: "Singapore", country: "SGN", dialingCode: 65, terminals: 2}

	airports := []airport{heathrowAirport, colomboAirport, changiAirport}
	fmt.Println(airports)

	airportsPointers := []*airport{&heathrowAirport, &colomboAirport, &changiAirport}
	fmt.Println(airportsPointers)

	floatArray := make([]float32, 100)
	floatArray[0] = 123;
	println("Float 32 index 0", floatArray[0])
	println(floatArray)
	fmt.Println("Using fmt print array", floatArray)

	mapExample := make(map[int] string)
	mapExample[2] = "Test Mals"
	fmt.Println(mapExample)
}
