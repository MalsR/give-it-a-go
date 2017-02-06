package main
//package name need to 'main' if main.go calls any functions in this file

import (
    "fmt"
)

/**
Simple method that demonstrates go functions with single and multiple return types, functions with named return types and 
annonymous functions
**/
func demonstrateGoFunctions() {	
	term, sum := numbersSum(2, 3, 4, 5)
	fmt.Println(term, sum)

	anotherTerm, anotherSum := sumWithNamedReturnTypes(2, 3, 4, 5, 6)
	fmt.Println(anotherTerm, anotherSum)
    
	addFunction := func(numbers ...int) (term int, sum int) {
		for _, number := range numbers {
			sum += number
			term += 1
		}
		return
	}
	numberTerm, numberSum := addFunction(33, 1, 2, 60)
	fmt.Println("Term is: ", numberTerm, "Sum is:", numberSum)
}

func numbersSum(numbers ...int) (int, int) {
	result := 0
	term := 0
	for _, number := range numbers {
		term += 1
		result += number
	}
	return term, result
}

func sumWithNamedReturnTypes(numbers ...int) (term int, result int) {
	//name the return types in method definition and just return
	for _, number := range numbers {
		result += number
		term += 1
	}
	return
}

