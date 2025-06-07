package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "G👨‍💻o"

	for i := 0; i < len(str); i++ {
		fmt.Println(reflect.TypeOf(str[i]))
	}

	for _, s := range str {
		fmt.Println(reflect.TypeOf(s))
		fmt.Printf("%v\n", s)
	}

}
