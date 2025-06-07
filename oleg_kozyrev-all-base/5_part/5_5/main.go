package main

import "fmt"

type SomeStruct struct {
	Value int
}

func CheckForNil(i interface{}) {
	if i == nil {
		fmt.Println("it is nil")
		return
	}

	fmt.Println("it is not nil")
}

func main() {
	var a *SomeStruct

	CheckForNil(a)
}
