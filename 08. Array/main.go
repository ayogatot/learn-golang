package main

import "fmt"

func main() {
	fmt.Println("============== Array ==============")
	fmt.Printf("\n")
	// Go array concept is similar with the other programming language
	// Array example :
	var thisArray [3]string
	thisArray[0] = "Monkey"
	thisArray[1] = "D."
	thisArray[2] = "Luffy"
	fmt.Println(thisArray[0], thisArray[1], thisArray[2])

	// Array first initiation
	// Horizontal Style
	var Array2 = [4]string{"Roronoa", "Zoro", "Sanji", "Vinsmoke"}
	fmt.Println("\nJumlah Array2 =>", len(Array2))
	fmt.Println("Isi Array2 =>", Array2)

	// Array first initiation
	// Vertical Style
	var Array3 = [3]string{
		"Monkey",
		"D.",
		"Dragon",
	}
	fmt.Println("\nArray3 =>", Array3)

	// Array first initiation without limit
	var Array4 = [...]string{"Monkey", "D.", "Grap"}
	fmt.Println("\nArray4 =>", Array4)

	// Array Multidimensi
	var Array5 = [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}
	fmt.Println("\nArray5 =>", Array5)
	fmt.Printf("\n")

	// Declare Array using Make
	var Array6 = make([]string, 2)
	Array6[0] = "Sanji"
	Array6[1] = "Vinsmoke"
	fmt.Println(Array6[0], Array6[1])
	fmt.Printf("\n")

	// Looping Array with for
	for i := 0; i < len(Array4); i++ {
		fmt.Printf("%s\n", Array4[i])
	}
	fmt.Printf("\n")

	// Looping Array with for-range
	for i, item := range Array4 {
		fmt.Printf("Elemen ke %d = %s\n", i, item)
	}
	fmt.Printf("\n")
	// alternative for-range
	for _, item := range Array3 {
		fmt.Printf("Isi Array3 : %s\n", item)
	}

	// If only indeks needed :
	// for i, _ := range Array3 {}
	// or
	// for i := range Array3 {}

	fmt.Println("\n============== Array ==============")

}
