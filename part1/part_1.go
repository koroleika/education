package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func equation(a, b, c float64) {
	var D = b*b - 4*a*c
	var sqrtD = math.Sqrt(D)

	if D > 0 {
		fmt.Printf("Equation has two roots\n")
		FirstRoot := (-b + sqrtD) / (2 * a)
		SecondRoot := (-b - sqrtD) / (2 * a)
		fmt.Printf("First root is :%f\n", FirstRoot)
		fmt.Printf("Second root is :%f\n", SecondRoot)

	} else if D == 0 {
		fmt.Printf("Equation has only one root\n")
		OneRoot := -b / (2 * a)
		fmt.Printf("This is:%f\n", OneRoot)
	} else {
		fmt.Printf("Equation has no roots\n")
	}

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var (
		FirstDigit    float64
		SecondDigit   float64
		ThirdDigit    float64
		OkOrNot       string
		FirstDigitNo  string
		SecondDigitNo string
		ThirdDigitNo  string
	)

	//fmt.Println("Hello world")
	//fmt.Printf("Name is: %s\nAge is: %d\n", name, age)
	/*
		fmt.Printf("Input first digit: ")
		fmt.Scan(&FirstDigit)
		fmt.Printf("Input second digit: ")
		fmt.Scan(&SecondDigit)
		fmt.Printf("Input third digit: ")
		fmt.Scan(&ThirdDigit)
	*/

	FirstDigit, err := strconv.ParseFloat(os.Args[1], 2)
	check(err)

	SecondDigit, err = strconv.ParseFloat(os.Args[2], 2)
	check(err)

	ThirdDigit, err = strconv.ParseFloat(os.Args[3], 2)
	check(err)

	switch {
	case SecondDigit > 0 && ThirdDigit > 0:
		fmt.Printf("You have an equation like this: %.2fx^2+%.2fx+%.2f = 0\n", FirstDigit, SecondDigit, ThirdDigit)
	case SecondDigit > 0:
		fmt.Printf("You have an equation like this: %.2fx^2+%.2fx%.2f = 0\n", FirstDigit, SecondDigit, ThirdDigit)
	case ThirdDigit > 0:
		fmt.Printf("You have an equation like this: %.2fx^2%.2fx+%.2f = 0\n", FirstDigit, SecondDigit, ThirdDigit)
	case SecondDigit < 0 && ThirdDigit < 0:
		fmt.Printf("You have an equation like this: %.2fx^2%.2fx%.2f = 0\n", FirstDigit, SecondDigit, ThirdDigit)
	}

	/*
		if SecondDigit > 0 && ThirdDigit > 0 {
			fmt.Printf("You have an equation like this: %.2fx^2+%.2fx+%.2f = 0\n", FirstDigit, SecondDigit, ThirdDigit)
		} else if SecondDigit > 0 {
			fmt.Printf("You have an equation like this: %.2fx^2+%.2fx%.2f = 0\n", FirstDigit, SecondDigit, ThirdDigit)
		} else if ThirdDigit > 0 {
			fmt.Printf("You have an equation like this: %.2fx^2%.2fx+%.2f = 0\n", FirstDigit, SecondDigit, ThirdDigit)
		} else {
			fmt.Printf("You have an equation like this: %.2fx^2%.2fx%.2f = 0\n", FirstDigit, SecondDigit, ThirdDigit)
		}
	*/
	fmt.Printf("Is it right? (y/n)	")
	fmt.Scan(&OkOrNot)
	OkOrNot = strings.ToLower(OkOrNot)
	/*
				pattern := `\d*.\d`
				Digit = 43
				Digit= strconv.Atoi(Digit)
				Matched, err := regexp.Compile(pattern, Digit)

			pattern := `\d*.\d`
			fmt.Fscan(&FirstDigit1 io.Reader , pattern)

		//r := string.NewReader("101010")
		n,err := fmt.Fscanf(FirstDigit1, "%d", &FirstDigit1)
		fmt.Println(n)
		pattern := `^\d*$|^\d*.\d*$|^\d*,\d*$`
	*/
	pattern := `^[\-\+\.]?\d{1,}$|^[\-\+]?\d{1,}[\.]?\d*$|^[\-\+]?[\.]{1}\d{1,}$`
	if OkOrNot == "y" || OkOrNot == "yes" {
		//equation(FirstDigitNo, SecondDigitNo, ThirdDigitNo)
	} else if OkOrNot == "n" || OkOrNot == "no" {

		fmt.Print("Input first digit: ")
		fmt.Scan(&FirstDigitNo)

		Matched, err := regexp.Match(pattern, []byte(FirstDigitNo))
		check(err)
		if Matched {
			FirstDigitNo, err = strconv.ParseFloat(os.Args[1:], 2)
			check(err)
		} else {
			for Matched != true {
				fmt.Println("Invalid input data")
				fmt.Print("Input first digit again: ")
				fmt.Scan(&FirstDigitNo)
				Matched, err = regexp.Match(pattern, []byte(FirstDigitNo))
				check(err)
			}
		}

		fmt.Printf("Input second digit: ")
		fmt.Scan(&SecondDigitNo)

		Matched, err = regexp.Match(pattern, []byte(SecondDigitNo))
		check(err)
		if Matched {
			fmt.Println("zaebis' 2")
		} else {
			fmt.Println("Invalid input data")
			fmt.Printf("Input second digit again: ")
			fmt.Scan(&SecondDigitNo)
		}

		fmt.Printf("Input third digit: ")
		fmt.Scan(&ThirdDigitNo)

		Matched, err = regexp.Match(pattern, []byte(ThirdDigitNo))
		check(err)
		if Matched {
			fmt.Println("zaebis' 3")
		} else {
			fmt.Println("Invalid input data")
			fmt.Printf("Input third digit again: ")
			fmt.Scan(&ThirdDigitNo)
		}

		//equation(FirstDigitNo, SecondDigitNo, ThirdDigitNo)
	}

	/*
		arg := os.Args[1:]
		fmt.Println(arg)

			Discrim = SecondDigit*SecondDigit - 4*FirstDigit*ThirdDigit
			SqrtDiscrim = math.Sqrt(Discrim)
			if Discrim > 0 {
				fmt.Printf("Equation has two roots\n")
				FirstRoot := (-SecondDigit + SqrtDiscrim) / (2 * FirstDigit)
				SecondRoot := (-SecondDigit - SqrtDiscrim) / (2 * FirstDigit)
				fmt.Printf("First root is :%f\n", FirstRoot)
				fmt.Printf("Second root is :%f\n", SecondRoot)

			} else if Discrim == 0 {
				fmt.Printf("Equation has only one root\n")
				OneRoot := -SecondDigit / (2 * FirstDigit)
				fmt.Printf("This is:%f\n", OneRoot)
			} else {
				fmt.Printf("Equation has no roots\n")
			}
	*/
}
