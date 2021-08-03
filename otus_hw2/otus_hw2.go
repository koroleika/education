package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile(`[\.\-,!?«»—]`)

func wordCounterSliceOnly(str *string) {

	strNew := re.ReplaceAllString(*str, "")
	strNew = strings.ToLower(strNew)

	//	wordCount := make(map[string]int)

	words := strings.Fields(strNew)
	/*	for i := 0; i < len(words); i++ {
			if _, ok := wordCount[words[i]]; ok {
				wordCount[words[i]] = wordCount[words[i]] + 1
			} else {
				wordCount[words[i]] = 1
			}
		}
	*/
	type keyValue struct {
		Key   string
		Value int
	}

	var sorted []keyValue
	for i := 0; i < len(words); i++ {

		sorted = append(sorted, keyValue{words[i], 1})
	}

	fmt.Println("---------------------------------------------")
	for i := 0; i < len(sorted); i++ {
		for j := i + 1; j < len(sorted); j++ {
			//fmt.Println(sorted[i].Key)
			//fmt.Println(sorted[j].Key)
			//fmt.Println(sorted[len(sorted)-1])
			//fmt.Println(i, j)
			if sorted[i].Key == sorted[j].Key {
				/*
					sorted[j] = sorted[len(sorted)-1]
					sorted[len(sorted)-1].Key = ""
					sorted[len(sorted)-1].Value = 0
					sorted = sorted[:len(sorted)-1]*/
				sorted = append(sorted[0:j], sorted[j:]...)
				sorted[i].Value = sorted[i].Value + 1
			}
		}
	}

	/*
		for i := 0; i < len(words); i++ {
			for j := 1; j < len(words); j++ {
				//fmt.Println(len(sorted))

				if sorted[j].Key == words[i] {
					sorted[j].Value = sorted[j].Value + 1
				} else {
					//fmt.Print("asd")
					sorted = append(sorted, keyValue{words[i], 1})
				}
			}
		}*/
	/*
		for key, value := range wordCount {
			sorted = append(sorted, keyValue{key, value})
		}
	*/

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	if len(sorted) < 10 {
		for _, keyValue := range sorted {
			fmt.Printf("The word ---%s--- occurs in the text %d times\n", keyValue.Key, keyValue.Value)
		}
	} else {
		for i := 0; i < 10; i++ {
			fmt.Printf("The word ---%s--- occurs in the text %d times\n", sorted[i].Key, sorted[i].Value)
		}
	}
}

func wordCounter(str *string) {

	strNew := re.ReplaceAllString(*str, "")
	strNew = strings.ToLower(strNew)

	wordCount := make(map[string]int)

	words := strings.Fields(strNew)
	for i := 0; i < len(words); i++ {
		if _, ok := wordCount[words[i]]; ok {
			wordCount[words[i]] = wordCount[words[i]] + 1
		} else {
			wordCount[words[i]] = 1
		}
	}

	type keyValue struct {
		Key   string
		Value int
	}

	var sorted []keyValue

	for key, value := range wordCount {
		sorted = append(sorted, keyValue{key, value})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	if len(sorted) < 10 {
		for _, keyValue := range sorted {
			fmt.Printf("The word ---%s--- occurs in the text %d times\n", keyValue.Key, keyValue.Value)
		}
	} else {
		for i := 0; i < 10; i++ {
			fmt.Printf("The word ---%s--- occurs in the text %d times\n", sorted[i].Key, sorted[i].Value)
		}
	}
	/*
		if len(sorted) < 10 {
			for _, keyValue := range sorted {
				if keyValue.Value > 1 {
					fmt.Printf("The word ---%s--- occurs in the text %d times\n", keyValue.Key, keyValue.Value)
				}
			}
		}
	*/
}

func check(err error) {
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}
}

func scan() {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	check(err)
	wordCounter(&str)
}

func scanFromFile() {
	file, err := ioutil.ReadFile("text2.txt")
	if err != nil {
		fmt.Println("read fail", err)
	}
	str := string(file)
	wordCounter(&str)
	fmt.Println("-----------------Другой метод-----------------")
	wordCounterSliceOnly(&str)
}

func main() {
	//scan()
	scanFromFile()
}
