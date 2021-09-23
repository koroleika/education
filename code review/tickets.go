package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, n int
	var arr []int
	c := 0
	fmt.Print("Введите количество билетов: ")
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil && a <= 0 {
		fmt.Println("Неправильное количество билетов, выходим из функции")
		os.Exit(1)
	}
	fmt.Print("Введите номера билетов (через Enter/точку/пробел): ")
	for g := 0; g < n; g++ {
		if g > 0 {
			fmt.Printf("Билет %d из %d : ", g+1, n)
		}
		t, err2 := fmt.Scanf("%d", &a)
		if err2 != nil || t != 1 {
			fmt.Println("Неправильный номер билета")
			g--
			fmt.Print("Введите номер еще раз: ")
			continue
		}
		if g == 0 {
			fmt.Printf("Билет %d из %d : ", g+1, n)
		}
		b = a
		arr = make([]int, 0)
		for a > 0 {
			arr = append(arr, a%10)
			a /= 10
		}
		if len(arr)%2 != 0 || len(arr) < 6 || len(arr) > 8 {
			fmt.Println("Допустимы 6-ти,8-ми,10-ти значные числа, введенный разряд числа:", len(arr))
			g--
			fmt.Print("Введите номер еще раз: ")
			continue
		} else {
			sum1 := 0
			sum2 := 0
			for i := 0; i < len(arr)/2; i++ {
				sum1 = sum1 + arr[i]
			}
			for i := len(arr) / 2; i < len(arr); i++ {
				sum2 = sum2 + arr[i]
			}
			if sum1 == sum2 {
				fmt.Print("Билет \"", b, "\" счастливый! Поздравляю!\n")
				c = c + 1
			} else {
				fmt.Print("Билет \"", b, "\" не счастливый! Не расстраивайтесь!\n")
			}
		}
	}
	fmt.Printf("Количество счастливых билетов: %d из %d!", c, n)
}