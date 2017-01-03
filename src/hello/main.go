package main

//It is a go convention to name the filename as main.go in a single file go application 
//all go programs must have a main package and main function which is the entry point for a go program

//package imports go after the package definitions, import the format package of go
import (
  "fmt"
  "os"
)

func main() {
  // len() built in function that gets the length of some input type
  fmt.Println("Lenth of arguments", len(os.Args))

  if len(os.Args) > 1 {
    // os.Args returns an array with the 'program arguments, including the name of the go file being executed followed
    //by any other user input
    fmt.Println("User input found:", os.Args)

  //  Example of what a command argument looks like when:  go run main.go "Hello Whats the plan?"
  //  os.Args[0] go run main.go
  //  os.Args[1] "Hello Whats the plan?"
    fmt.Println(os.Args[0])
    fmt.Println(os.Args[1])

  } else {
    fmt.Println("Planes Rule")
  }
}

//to compile the project do 'go build' on root project, in this case 'hello'. This will create 
//an extension less file hello (run ./hello)
//to run go programs similar to running like a python 'go run main.go'

//'gofmt main.go' to format go source code
// run 'gofmt -w main.go' on a file to format the file and write and update source file as opposed to standard output i.r. '-w'
