package main

import "fmt"

func main() {
	a1 := make([]int, 0, 10)

	fmt.Printf("len=%d, cap=%d\n", len(a1), cap(a1))

	a1 = append(a1, []int{1, 2, 3, 4, 5}...)
	//a1 := []int{1, 2, 3, 4, 5}

	fmt.Printf("ptr=%p, len=%d, cap=%d\n", &a1[0], len(a1), cap(a1))

	a2 := append(a1, 6)

	fmt.Printf("ptr=%p, len=%d, cap=%d\n", &a2[0], len(a2), cap(a2))

	a3 := append(a1, 7)

	fmt.Printf("ptr=%p, len=%d, cap=%d\n", &a3[0], len(a3), cap(a3))

	fmt.Println(a1, a2, a3)
}
