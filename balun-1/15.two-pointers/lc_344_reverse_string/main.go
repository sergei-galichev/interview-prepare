package main

import "fmt"

func main() {
	str := "hello"

	a := []byte(str)
	reverseString(a)

	fmt.Println(string(a))

}

func reverseString(s []byte) {
	first := 0
	second := len(s) - 1

	for first < second {
		s[first], s[second] = s[second], s[first]

		first++
		second--
	}
}
