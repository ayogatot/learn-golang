package main

import (
	"fmt"
	"math"
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

	fmt.Printf("\n%d\n", divideNumber(6, 3))
	fmt.Printf("%d\n", divideNumber(2, 1))
	fmt.Printf("%d\n", divideNumber(5, 0))

	var diameter float64 = 15
	var area, circumference = calculate(diameter)
	var area2, circumference2 = calculate(diameter)

	fmt.Printf("\nluas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

	fmt.Printf("\nluas lingkaran\t\t: %.2f \n", area2)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference2)

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

// Functions with multiple return
func calculate(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}

func calculated(d float64) (circumference, area float64) {
	area = math.Phi * math.Pow(d/2, 2)
	circumference = math.Phi * d

	return
}
