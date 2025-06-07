package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "ddЯЙ福"

	fmt.Println("Длина через 'len':", len(str))
	fmt.Println("Длина через 'RuneCountInString':", utf8.RuneCountInString(str))

}
