// Вычисление факториала
package part4

//package main

import (
	"fmt"
	"math/big"
	"time"
)

// создание структуры, дял хранения простых чисел и степени для функции factDegree
type degreeDigit struct {
	Digit  int
	Degree int
}

// обычная рекурсия
func factRec(n int) *big.Int {
	result := big.NewInt(1)
	for n != 0 {
		result.Mul(result, big.NewInt(int64(n)))
		n = n - 1
		factRec(n)
	}
	return result
}

// поиск факториала с помощью цикла while
func factCycleWhile(n int) *big.Int {
	result := big.NewInt(1)
	for n > 1 {
		result.Mul(result, big.NewInt(int64(n)))
		n--
	}
	return result
}

// поиск факториала с помощью цикла for
func factCycleFor(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

// поиск факториала с помощью дерева
func prodTree(l, r int) *big.Int {
	result := big.NewInt(1)
	if l > r {
		return big.NewInt(1)
	}
	if l == r {
		return big.NewInt(int64(l))
	}
	if r-l == 1 {
		return big.NewInt(int64(l * r))
	}
	m := (l + r) / 2
	result.Mul(prodTree(l, m), prodTree(m+1, r))
	return result
}

func factTree(n int) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	}
	if n == 0 {
		return big.NewInt(1)
	}
	if n == 1 || n == 2 {
		return big.NewInt(int64(n))
	}
	return prodTree(2, n)
}

// поиск факториала с помощью разложения числа N на простые числа и подсчета количеств их вхождений в число N
func factDegree(n int) *big.Int {

	var degreeD []degreeDigit
	seive := make([]bool, n+1)

	for i := 2; i < n; i++ {
		if !seive[i] {
			k := n / i
			c := 0
			for k > 0 {
				c += k
				k /= i
			}
			degreeD = append(degreeD, degreeDigit{i, c})
			j := 2
			for (i * j) <= n {
				seive[i*j] = true
				j++
			}
		}
	}

	result := big.NewInt(1)
	de := big.NewInt(1)
	for _, degreeDigit := range degreeD {
		de.Exp(big.NewInt(int64(degreeDigit.Digit)), big.NewInt(int64(degreeDigit.Degree)), nil)
		result.Mul(de, result)
	}
	return result
}

func main() {
	start := time.Now()
	fmt.Println("Recursion: ", factRec(30))
	duration := time.Since(start)
	fmt.Println("Work time recursion: ", duration)

	start = time.Now()
	fmt.Println("Cycle while: ", factCycleWhile(100))
	duration1 := time.Since(start)
	fmt.Println("Work time cycle while: ", duration1.Nanoseconds())

	start = time.Now()
	fmt.Println("Cycle for: ", factCycleFor(100))
	duration2 := time.Since(start)
	fmt.Println("Work time cycle for: ", duration2.Nanoseconds())

	start = time.Now()
	fmt.Println("Tree func: ", factTree(10))
	duration3 := time.Since(start)
	fmt.Println("Work time Tree func: ", duration3.Nanoseconds())

	start = time.Now()
	fmt.Println("Degree func: ", factDegree(10))
	duration4 := time.Since(start)
	fmt.Println("Work time Degree func: ", duration4.Nanoseconds())
}
