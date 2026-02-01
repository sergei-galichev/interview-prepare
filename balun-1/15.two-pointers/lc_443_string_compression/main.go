package main

import (
	"fmt"
	"strconv"
)

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}

	length := compress(chars)

	fmt.Println(string(chars))
	fmt.Println(length)
}

func compress(chars []byte) int {
	idx1 := 0
	idx2 := 0

	for idx2 < len(chars) {
		groupSize := 1

		for groupSize+idx2 < len(chars) && chars[idx2] == chars[idx2+groupSize] {
			groupSize++
		}

		chars[idx1] = chars[idx2]
		idx1++

		if groupSize > 1 {
			for _, digit := range strconv.Itoa(groupSize) {
				chars[idx1] = byte(digit)

				idx1++
			}
		}

		idx2 += groupSize
	}

	return idx1
}
