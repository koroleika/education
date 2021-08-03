/*

Распаковка строки
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:

* "a4bc2d5e" => "aaaabccddddde"
* "abcd" => "abcd"
* "45" => "" (некорректная строка)

Дополнительное задание: поддержка escape - последовательности
* `qwe\4\5` => `qwe45` (*)
* `qwe\45` => `qwe44444` (*)
* `qwe\\5` => `qwe\\\\\` (*)

*/
package main

import (
	"fmt"
)

//str[i] == '\\' ||
func unpackEsc(str string) {
	var strResult string
	for i := 0; i < len(str); i++ {
		if (str[i] == '0' || str[i] == '1' || str[i] == '2' || str[i] == '3' || str[i] == '4' ||
			str[i] == '5' || str[i] == '6' || str[i] == '7' || str[i] == '8' || str[i] == '9') && str[i-1] != '\\' {
			n := int(str[i] - '0')
			//fmt.Println(n)
			for j := 0; j < n-1; j++ {
				strResult = strResult + string(str[i-1])
			}
		} else {
			if str[i] == '\\' && str[i-1] == '\\' && (str[i+1] == '0' || str[i+1] == '1' ||
				str[i+1] == '2' || str[i+1] == '3' || str[i+1] == '4' ||
				str[i+1] == '5' || str[i+1] == '6' || str[i+1] == '7' || str[i+1] == '8' || str[i+1] == '9') {
				n := int(str[i+1] - '0')
				//fmt.Println(n)
				for j := 0; j < n; j++ {
					strResult = strResult + string(str[i])
				}
				i++
			} else {
				if str[i] == '\\' {

				} else {
					strResult = strResult + string(str[i])
				}
			}

		}
		//if str[i] == '\\' && str[i+1] == '4'{

		//}
	}
	fmt.Print("Unpacked string: ", strResult)

}

func unpack(str string) {
	var strResult string
	for i := 0; i < len(str); i++ {
		if str[i] == '0' || str[i] == '1' || str[i] == '2' || str[i] == '3' || str[i] == '4' ||
			str[i] == '5' || str[i] == '6' || str[i] == '7' || str[i] == '8' || str[i] == '9' {
			n := int(str[i] - '0')
			//fmt.Println(n)
			for j := 0; j < n-1; j++ {
				strResult = strResult + string(str[i-1])
			}
		} else {
			strResult = strResult + string(str[i])
		}
	}
	fmt.Print("Unpacked string: ", strResult)
}

func main() {
	var (
		str string
	)
	fmt.Print("Input string: ")
	fmt.Scan(&str)
	unpackEsc(str)
}
