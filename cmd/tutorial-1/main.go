package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello World")

	var myNum int = 2000
	var myNum2 int32 = 1000
	myNum3 := 2
	result := myNum * int(myNum2) / myNum3
	for value := range result {
		if value == 500000 {
			fmt.Println(value)
		}
	}

	var s string = "résumé"         // string = sequence of bytes
	var r []rune = []rune("résumé") // rune = single int32. Here is slice of int32's
	sbytes := []int16{}
	fmt.Println("String Bytes:")
	for value := range len(s) {
		sbytes = append(sbytes, int16(s[value]))
	}
	fmt.Println(sbytes)
	rbytes := []int16{}
	fmt.Println("Rune Bytes:")
	for value := range len(r) {
		rbytes = append(rbytes, int16(r[value]))
	}
	fmt.Println(rbytes)

	fmt.Println(sumdigits(123))
	fmt.Println(sumdigits(457))
	fmt.Println(sumdigits(32))
}

func sumdigits(input int) int {
	str := strconv.Itoa(input)
	arr := strings.Split(str, "")
	value := 0
	for _, element := range arr {
		element, _ := strconv.Atoi(element)
		value += int(element)
	}
	return value
}
