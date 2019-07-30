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

	fmt.Println("\n============== Function ==============")

}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}
