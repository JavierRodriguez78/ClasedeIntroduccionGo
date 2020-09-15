package main

import "fmt"

func suma(a int, b int) int {
	return a + b
}

// func rectPer(length, width float64) (float64, float64) {
// 	var area = length * width
// 	var perimeter = (length + width) * 2
// 	return area, perimeter
// }

func rectangle(l int, b int) (area int) {
	var parameter int
	parameter = 2 * (l + b)
	fmt.Println("Parameter: ", parameter)
	area = l * b
	return
}

func rectPer(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func variadicEx(number int, str ...string) {
	fmt.Println(str[0])
	fmt.Println(str[3])
}
