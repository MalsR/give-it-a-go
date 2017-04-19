package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	userArgs := os.Args
	fmt.Println("Simple CLI greetings")
	fmt.Println(userArgs)

	hourOfDay := time.Now().Hour()
	message, timeError := printHourOfDay(hourOfDay)
	if timeError != nil {
		fmt.Println(timeError)
		os.Exit(1)
	}

	fmt.Println(message)

}

//can go return multiple values?
func printHourOfDay(hourOfDay int) (string, error) {
	var message string
	fmt.Println("Hour of day is", hourOfDay)

	//just to demonstrate a go error
	if hourOfDay == 23 {
		timeError := errors.New("Too late to GO... get it?")
		//message has no value, in go "zero value" is the default value for a data type, in this case "" for string message
		return message, timeError
	}

	//nil value tells the calling function that no errors were encountered during execution of this function
	if hourOfDay < 12 {
		return "A very good morning to you folks", nil
	} else if hourOfDay > 12 && hourOfDay < 19 {
		return "It's that time of the evening", nil
	} else if hourOfDay > 19 && hourOfDay < 1 {
		return "Time to hit the bed", nil
	} else {
		return message, errors.New("BAM")
	}
}

func exampleMethodReturningMultipleValues() (string, string, error) {
	return "Hello", "Second Param", nil
}
