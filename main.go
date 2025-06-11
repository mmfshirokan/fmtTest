package main

import (
	"fmt"
)

type MyError struct {
	Info string
	Code int
}

func (e MyError) Error() string {
	return e.Info
}

var errMy error = MyError{
	Info: "Something went wrong",
	Code: 404,
}

var errEmpty error = MyError{
	Info: "",
	Code: 404,
}

func main() {
	fmt.Printf("v: %v\n", errMy)
	fmt.Printf("+v: %+v\n", errMy)
	fmt.Printf("s: %s\n", errMy)

	fmt.Printf("v: %v\n", errMy.(MyError))
	fmt.Printf("+v: %+v\n", errMy.(MyError))
	fmt.Printf("s: %s\n", errMy.(MyError))
}
