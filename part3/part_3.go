//Подсчет счастливых билетов
// написать алгоритм с рекурсией
package part

import (
	"fmt"
	"math/big"
	"time"
)

var Count, N int

func fastMethodRec(nr, sum1, sum2 int) {
	if nr == 2*N {
		if sum1 == sum2 {
			Count++
		}
		return
	}
	for i := 0; i <= 9; i++ {
		if nr < N {
			fastMethodRec(nr+1, sum1+i, sum2)
		} else {
			fastMethodRec(nr+1, sum1, sum2+i)
		}
	}
}

func fastMethod9Cycles() {
	var count int
	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for c := 0; c < 10; c++ {
				for d := 0; d < 10; d++ {
					for e := 0; e < 10; e++ {
						for f := 0; f < 10; f++ {
							for g := 0; g < 10; g++ {
								for h := 0; h < 10; h++ {
									for i := 0; i < 10; i++ {
										j := a + b + c + d + e - f - g - h - i
										if j >= 0 && j <= 9 {
											count++
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("\nAnswer from fast way is(9 cycles): ", count)
}

func slowMethod10Cycles() {
	var count int
	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for c := 0; c < 10; c++ {
				for d := 0; d < 10; d++ {
					for e := 0; e < 10; e++ {
						for f := 0; f < 10; f++ {
							for g := 0; g < 10; g++ {
								for h := 0; h < 10; h++ {
									for i := 0; i < 10; i++ {
										for j := 0; j < 10; j++ {
											if a+b+c+d+e == f+g+h+i+j {
												count++
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("\nAnswer from slow way is(10 cycles): ", count)
}

func fastMethod7Cycles() {
	var count int
	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for c := 0; c < 10; c++ {
				for d := 0; d < 10; d++ {
					for e := 0; e < 10; e++ {
						for f := 0; f < 10; f++ {
							for g := 0; g < 10; g++ {
								h := a + b + c + d - e - f - g
								if h >= 0 && h <= 9 {
									count++
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("\nAnswer from fast way is(7 cycles): ", count)
}

func fastMethod5Cycles() {
	var count int
	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for c := 0; c < 10; c++ {
				for d := 0; d < 10; d++ {
					for e := 0; e < 10; e++ {
						f := a + b + c - d - e
						if f >= 0 && f <= 9 {
							count++
						}
					}
				}
			}
		}
	}
	fmt.Println("\nAnswer from fast way is(5 cycles): ", count)
}

// Медленный метод 6 циклов
func slowMethod() {
	var count int
	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for c := 0; c < 10; c++ {
				for d := 0; d < 10; d++ {
					for e := 0; e < 10; e++ {
						for f := 0; f < 10; f++ {
							if a+b+c == d+e+f {
								count++
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("\nAnswer from slow way is(6 cycles): ", count)
}

func slowMethod8Cycles() {
	var count int
	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for c := 0; c < 10; c++ {
				for d := 0; d < 10; d++ {
					for e := 0; e < 10; e++ {
						for f := 0; f < 10; f++ {
							for g := 0; g < 10; g++ {
								for h := 0; h < 10; h++ {
									if a+b+c+d == e+f+g+h {
										count++
									}
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("\nAnswer from slow way is(8 cycles): ", count)
}

func fastMethod2Cycles() {
	var count int
	for a := 0; a < 1000; a++ {
		for b := 0; b < 1000; b++ {
			if (a/100 + a/10%10 + a%10) == (b/100 + b/10%10 + b%10) {
				count++
			}
		}
	}
	fmt.Println("\nAnswer from fast way is(2 cycles)", count)
}

func fastMethod1Cycle() {
	var count int
	for a := 0; a < 1000000; a++ {
		if (a/100000 + a/10000%10 + a/1000%10) == (a/100%10 + a/10%10 + a%10) {
			count++
		}

	}
	fmt.Println("\nAnswer from fast way is(1 cycle)", count)
}

func fuc(n int) *big.Int {
	result := big.NewInt(1)
	for n != 0 {
		if n >= 60 {
			break
		}
		result.Mul(result, big.NewInt(int64(n)))
		n = n - 1
		fuc(n)
	}
	return result
	//return(n*fuc(n-1))
	// Найти способ экстренного выхода из рекурсии
	// Оптимизировать алгоритм подсчета
}

func main() {
	start := time.Now()
	slowMethod()
	duration := time.Since(start)
	fmt.Println("Work time slow way(6 cycles): ", duration)

	start = time.Now()
	fastMethod5Cycles()
	duration4 := time.Since(start)
	fmt.Println("Work time fast way(5 cycles): ", duration4)

	start = time.Now()
	fastMethod2Cycles()
	duration1 := time.Since(start)
	fmt.Println("Work time fast way(2 cycles): ", duration1)

	start = time.Now()
	fastMethod1Cycle()
	duration2 := time.Since(start)
	fmt.Println("Work time fast way(1 cycles): ", duration2)

	start = time.Now()
	slowMethod8Cycles()
	duration3 := time.Since(start)
	fmt.Println("Work time slow way(8 cycles): ", duration3)

	start = time.Now()
	fastMethod7Cycles()
	duration5 := time.Since(start)
	fmt.Println("Work time slow way(7 cycles): ", duration5)

	start = time.Now()
	slowMethod10Cycles()
	duration7 := time.Since(start)
	fmt.Println("Work time slow way(10 cycles): ", duration7)

	start = time.Now()
	fastMethod9Cycles()
	duration6 := time.Since(start)
	fmt.Println("Work time fast way(9 cycles): ", duration6)

	//fmt.Println(fuc(400))
	start = time.Now()
	N = 5
	fastMethodRec(0, 0, 0)
	fmt.Println(Count)
	duration8 := time.Since(start)
	fmt.Println("Work time REC way(10 digit tickets): ", duration8)
}

// 1000 000
/*
999 999
999 998

1234 5678
для 6ти значных билетов минус один цикл
написать все с рекурсией


*/
