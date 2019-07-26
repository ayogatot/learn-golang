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
	fmt.Println(Array5[0][0])

	fmt.Println("\n============== Array ==============")

}
