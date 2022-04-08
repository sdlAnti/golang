//Форматированный вывод
package main

import "fmt"

var a float64 = 5566.54563
var n1 byte = 'w'
var n2 int = 123

func main() {
	fmt.Printf("this is %f, type is %T", a, a)
	fmt.Printf("\nin quotes - %q; \nbinary mode - %b;\n", n1, n2)
	fmt.Println(a)
}
