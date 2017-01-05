package main

import (
	"fmt"
	"os"
	"time"
	"errors"
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

	if hourOfDay == 23 {
		timeError := errors.New("Too late to GO... get it?")
		//message has no value, in go "zero value" is the default value for a data type, in this case "" for string message
		return message, timeError
	}

	if hourOfDay < 12 {

	} else if hourOfDay > 12 && hourOfDay < 21 {

	}
	//nil value tells the calling function that no errors were encountered during execution of this function
	return "No Errors - Hour of day is", nil
}
