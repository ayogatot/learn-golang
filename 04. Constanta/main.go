package main

import "fmt"

func main() {
	fmt.Println("Constanta")
	// variable that cannot be changing after first initiation
	const thisNumber, thisString, thisBool, thisFloat = 192, "Ya ini kan API", true, 129019.1281289
	fmt.Println(thisNumber, thisBool, thisFloat, thisString)

	// declaring variable using const can't use :=
	const numb string = "ohyeah"
	const pi = 3.14
	const jomblo = true
	fmt.Printf("Kamu itu %.3f dan pi itu %t dan numb itu %s \n", pi, jomblo, numb)
}
