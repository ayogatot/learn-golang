package main

import "fmt"

func main() {
	// all variable must be used
	// manifest typing = data type when declare variable must to be write
	var number int
	number = 10
	fmt.Printf("Number : %d !\n", number)

	// declaring variable without data type
	var i, j = 1, 2
	fmt.Println("Number2 :", i, j)

	// using := to replace var
	k := "OH YEAH"
	fmt.Println(k)

	// declaring variable using block
	var (
		l int    = 10
		s string = "OKE"
	)
	fmt.Println(l, s)

	// type inference to declare multiple variable
	thisNumber, thisString, thisBool, thisFloat := 1, "Ini String", true, 2.2
	fmt.Println(thisNumber, thisString, thisBool, thisFloat)

	// predefined variable (_)
	_ = "Ini variable sampah katanya"

	// declare variable using new
	iniString := new(string)
	// pointer string to showing memory address
	fmt.Println(iniString)
	// to show the true form of iniString
	fmt.Println(*iniString)

}
