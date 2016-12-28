package main

//It is a go convention to name the filename as main.go in a single file go application 
//all go programs must have a main package and main function which is the entry point for a go program

//package imports go after the package definitions, import the format package of go
import "fmt" 

func main() {
  fmt.Println("Planes Rule")
}

//to compile the project do 'go build' on root project, in this case 'hello'. This will create 
//an extention less file hello (run ./hello)
//to run go programs similar to running like a python 'go run main.go'

//'gofmt main.go' to format go source code
// run 'gofmt -w main.go' on a file to format the file and write and update source file as opposed to standard output i.r. '-w'
