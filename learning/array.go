package main

import "fmt"

func main() {
	var a [3]int
	a = [3]int{1, 2, 3}
	var b [3]int = [3]int{1, 2, 3}
	c := [3]int{1, 2, 3}
	d := [...]int{1, 2, 3, 4, 5, 6}
	e := [3]int{1: 12, 0: 7}

	fmt.Printf("%v\n", a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}
