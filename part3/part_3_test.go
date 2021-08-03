package part

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
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
