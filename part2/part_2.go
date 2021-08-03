// Возведения числа в степень(быстрый алгоритм)
package main

import (
	"fmt"
	"math/big"
	"time"
)

// Функция возведения числа в степень
// Принимает на вход три параметра:
// число, которое возводится в степень
// номер последнего значащего бита
// степень, в двоичном представлении
func Degree(digit, m int, bin [32]int) {
	// Результирующая переменная типа big.Int
	result := big.NewInt(1)
	// Идем циклом до последнего значащего бита
	for i := 0; i <= m; i++ {
		// Возведение 2 в степень i
		twoInDegree := 1 << i
		// Возведения числа в степень 2 в степени i
		a := new(big.Int).Exp(big.NewInt(int64(digit)), big.NewInt(int64(twoInDegree)), nil)
		// Возведение числа а в степень равную i-ому биту
		b := (new(big.Int).Exp(a, big.NewInt(int64(bin[i])), nil))
		// Произведение результатов
		result.Mul(result, b)
	}
	// Вывод ответа
	fmt.Print("Fast way answer is: ")
	resultString := result.String()
	fmt.Println(resultString[(len(resultString) - 4):(len(resultString))])
}

// Функция для перевода в систему счисления с основанием 2
// В функцию передается два параметра: число, которое возводится в степень и степень, которая и будет
// переводится в двоичную систему счисления
func ToBinary(digit, x int) {
	// Массив целых чисел размерносью 32 для представления степени в двоичной системе счисления
	var buff [32]int
	for i := 31; i >= 0; i-- {
		// Маска с которой будет производится логиеское & определяется сдвигом влево на i бит
		mask := 1 << i
		// Если получившееся число в результате не равно 0, то в массив записываем 1
		// Если получившееся число в результате равно 0, то в массив записываем 0
		if (x & mask) != 0 {
			buff[i] = 1
		} else {
			buff[i] = 0
		}
	}
	// Отсекаем не значащие 0 получившейся двоичной записи степени
	for j := 31; j >= 0; j-- {
		if buff[j] != 0 {
			// Вызываем функцию возведения в степень, передавая число, которое возводится в степень
			// номер последнего значащего бита
			// степень, в двоичном представлении
			Degree(digit, j, buff)
			// Выходим из цикла
			break
		}
	}
}

func main() {
	// Обявление переменных
	var (
		digit  int
		degree int
	)
	// Ввод переменных
	fmt.Print("Input digit: ")
	fmt.Scan(&digit)
	fmt.Print("Input degree: ")
	fmt.Scan(&degree)

	// Медленный способ возведения в степень
	start := time.Now()
	resultSlow := new(big.Int).Exp(big.NewInt(int64(digit)), big.NewInt(int64(degree)), nil)
	resultSlowString := resultSlow.String()
	fmt.Print("Slow way answer is: ")
	fmt.Println(resultSlowString[(len(resultSlowString) - 4):(len(resultSlowString))])
	duration1 := time.Since(start)
	fmt.Println("Work time slow way: ", duration1)

	// Вызов функции, для перевода в систему счисления с основанием 2(внутри вложена еще одна функция)
	start = time.Now()
	ToBinary(digit, degree)
	duration := time.Since(start)
	fmt.Println("Work time fast way: ", duration)
}
