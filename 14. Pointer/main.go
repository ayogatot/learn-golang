package main

import "fmt"

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	fmt.Println("")

	numberA = 5

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)
}

/*
	Ada dua hal penting yang perlu diketahui mengenai pointer:

	1. Variabel biasa bisa diambil nilai pointernya,
	   caranya dengan menambahkan tanda ampersand (&) tepat sebelum nama variabel. Metode ini disebut dengan referencing.
	2. Dan sebaliknya, nilai asli variabel pointer juga bisa diambil,
	   dengan cara menambahkan tanda asterisk (*) tepat sebelum nama variabel. Metode ini disebut dengan dereferencing.
*/
