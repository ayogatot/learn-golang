package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("============== Function ==============")
	fmt.Printf("\n")

	var names = []string{"John", "Wick"}
	printMessage("halo", names)

	a := 20.4
	t := 5.1
	fmt.Printf("Luas segitiga dengan alas %.2f dan tinggi %.2f adalah %.2f\n", a, t, luasSegitiga(a, t))

	fmt.Printf("%d\n", divideNumber(6, 3))
	fmt.Printf("%d\n", divideNumber(2, 1))
	fmt.Printf("%d\n", divideNumber(5, 0))

	fmt.Println("\n============== Function ==============")

}

// Function without return value
func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

// Function with return value
// if have return value, return value must be declared in functions
func luasSegitiga(alas, tinggi float64) float64 {
	result := alas * tinggi * 0.5
	return result
}

// Use return to terminate the functions
func divideNumber(a, b int) int {
	if b == 0 {
		fmt.Printf("%d cannot divided by ", a)
		return b
	}
	return a / b
}
