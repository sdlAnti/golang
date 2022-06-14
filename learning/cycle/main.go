/*
Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются. Числа в пределах от 0 до 10000.
Программа должна вывести цифры, которые имеются в обоих числах, через пробел. Цифры выводятся в порядке их нахождения в первом числе.
*/
package main

import "fmt"

func main() {
	var dig1, dig1reverse, dig2save, dig2, diga, digb int
	fmt.Scan(&dig1, &dig2)
	dig2save = dig2
	dig1reverse = 0
	if dig1 <= 10000 && dig2 <= 10000 {
		for ; dig1 > 0; dig1 /= 10 {
			dig1reverse = dig1reverse*10 + dig1%10
		}
		for ; dig1reverse > 0; dig1reverse /= 10 {
			diga = dig1reverse % 10
			dig2 = dig2save
			for ; dig2 > 0; dig2 /= 10 {
				digb = dig2 % 10
				if diga == digb {
					fmt.Print(diga, " ")
				}
			}
		}
	} else {
		fmt.Println("Ведено число больше 10000")
	}
}
