/*
Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются. Числа в пределах от 0 до 10000.
Программа должна вывести цифры, которые имеются в обоих числах, через пробел. Цифры выводятся в порядке их нахождения в первом числе.
*/
package main

import "fmt"

func main() {
	var dig1, dig2, diga, digb, dig2save int
	fmt.Scan(&dig1, &dig2)
	dig2save = dig2
	if dig1 <= 10000 && dig2 <= 10000 {
		for dig1 > 0 {
			diga = dig1 % 10
			dig1 /= 10
			dig2 = dig2save
			//fmt.Println("цикл 1 = ", diga)
			for dig2 > 0 {
				digb = dig2 % 10
				dig2 /= 10
				//	fmt.Println("цикл 2 = ", digb)
				if diga == digb {
					fmt.Print(diga, " ")
				}
			}
		}
	} else {
		fmt.Println("Ведено число больше 10000")
	}

}
