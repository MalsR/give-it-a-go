package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	userArgs := os.Args
	fmt.Println("Simple CLI greetings")
	fmt.Println(userArgs)
	hourOfDay := time.Now().Hour()
	printHourOfDay(hourOfDay)

}
func printHourOfDay(hourOfDay int) string {
	fmt.Println("Hour of day is", hourOfDay)
	if hourOfDay < 12 {

	} else if hourOfDay > 12 && hourOfDay < 21 {

	}
	return "Hour of day is"
}
