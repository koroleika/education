package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func Task(tasks []func() error, taskCount int, maxErrors int) {

	var start = make(chan int)
	var chErr = make(chan error, len(tasks))

	last := 0
	gorutines := false

	for i := 0; i < len(tasks); i++ {
		if i <= taskCount {
			go func(i int) {
				<- start
				chErr <- tasks[i]()
			}(i)
		} else {
			last = i
			gorutines = true
			break
		}
	}

	if gorutines && last < len(tasks) {
		go func() {
			<- start
			for i := last; i < len(tasks); i++ {
				chErr <- tasks[i]()
			}
		}()
	}
	close(start)

	countFunc := 0
	errorCount := 0
	for {
		err, _ := <- chErr
		countFunc++
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			errorCount++
			if errorCount > maxErrors {
				fmt.Println( "Колиество функций завершившихся с ошибкой: ", errorCount)
				os.Exit(1)
			}
		}
		if countFunc == len(tasks) {
			fmt.Println("Большинство функций завершилось без ошибок:", countFunc)
			break
		}
	}
}

func main() {
	funcSlice := []func() error{
		func() error {
			time.Sleep(1 * time.Second)
			println("1 func")
			return nil
		},
		func() error {
			time.Sleep(1 * time.Second)
			println("2 func")
			return errors.New("2 func is failed")
		},
		func() error {
			time.Sleep(1 * time.Second)
			println("3 func")
			return nil
		},
		func() error {
			time.Sleep(1 * time.Second)
			println("4 func")
			return errors.New("4 func is failed")
		},
		func() error {
			time.Sleep(1 * time.Second)
			println("5 func")
			return nil
		},
		func() error {
			time.Sleep(1 * time.Second)
			println("6 func")
			return errors.New("6 func is failed")
		},
		func() error {
			time.Sleep(1 * time.Second)
			println("7 func")
			return nil
		},
	}
	Task(funcSlice, 6, 0)
}