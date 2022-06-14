package main

import "fmt"

func main() {
	var b, c int = 31
	fmt.Println(b, c)

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]

	for elem := range a {
		elem = 100
		fmt.Println(elem)
	}
	fmt.Println(a) // [1 2 3 4 5]

	for idx := range a {
		fmt.Println(idx)
		a[idx] = 100
		fmt.Println(a[idx])

	}
	fmt.Println(a) // [100 100 100 100 100]
}
