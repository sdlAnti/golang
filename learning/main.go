package main

import "fmt"

func main() {
	var dig1, dig2, diga int
	fmt.Scan(&dig1, &dig2)
	if dig1 <= 10000 && dig2 <= 10000 {
		for dig1 > 0 {
			diga = dig1 % 10
			dig1 /= 10
			fmt.Println(diga)
			for i := 0; i < 5; i++ {
				fmt.Println("итерация", diga, "=", i)

			}
		}
	} else {
		fmt.Println("Ведено число больше 10000")
	}

}
