package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println(uniqN(10))
}

func uniqN(n int) []int {
	// in real projects use crypto/rand

	rand.Int()
	m := make(map[int]struct{}, n)
	res := make([]int, 0, n)

	for len(res) < n {
		tmp := rand.IntN(10)

		if _, ok := m[tmp]; !ok {
			res = append(res, tmp)
			m[tmp] = struct{}{}
		}
	}

	return res
}
