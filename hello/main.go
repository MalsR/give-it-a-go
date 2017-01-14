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

  //standard fo loop with init, condition and post statement, and the statements are optional says go
  for i := 0; i < 5; i++ {
    fmt.Println("Print ", i)
  }

  for 0 < 5 {
    fmt.Println("It's true")
    //Else the condition will always be true
    break
  }

  var numberToIterate = 0
  var condition = true
  for condition {
    fmt.Println(numberToIterate)
    numberToIterate++
    if numberToIterate > 5 {
      condition = false
    }
  }

  for {
    fmt.Println("For statement without any conditions.. yes")
    break
  }

  var stringArray [3]string
  stringArray[0] = "Index 0"
  stringArray[1] = "Index 1"
  stringArray[2] = "Index 2"
  //stringArray[3] = "Index 2" will result in src/hello/main.go:61: invalid array index 3 (out of bounds for 3-element array)
  fmt.Println(stringArray)

  //use slices, could say builtin on top of arrays, as it needs to grow new underlying array is created
  var stringSliceArray []string
  stringSliceArray = append(stringSliceArray, "Slice 0")
  stringSliceArray = append(stringSliceArray, "Slice 2")
  fmt.Println(stringSliceArray)

  //use range to iterate over the array, can also use fmt.Println(len(stringSliceArray)) to get size
  //the '_' character tells the go compiler that we will not be using a variable anywhere else
  for _, element := range stringSliceArray {
    fmt.Println(element)
  }
}

//to compile the project do 'go build' on root project, in this case 'hello'. This will create 
//an extension less file hello (run ./hello)
//to run go programs similar to running like a python 'go run main.go'

//'gofmt main.go' to format go source code
// run 'gofmt -w main.go' on a file to format the file and write and update source file as opposed to standard output i.r. '-w'

//goimports -w main.go will add missing imports and write to file and run e.g. if "os" package import is missing
//TODO do example with default go data types and default values