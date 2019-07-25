package main

import "fmt"

func main() {
	fmt.Println("============== Looping ==============")
	fmt.Println("")
	// Looping
	// Can be use as While or for each or for itself

	// normal for
	for i := 0; i < 5; i++ {
		fmt.Print(i, "\n")
	}

	fmt.Printf("\n")

	// for like while
	var index = 0

	for index < 5 {
		fmt.Println("Ini Angka", index)
		index++
	}

	// for without condition
	// infinite looping
	// can be stop by using break statement
	fmt.Printf("\n")
	var j = 0
	for {
		fmt.Println("For yang ke-", j)
		j++
		if j == 5 {
			fmt.Printf("yang ke-%d BREAK\n", j)
			break
		}
	}

	// Using continue statement to force to the next loop
	fmt.Printf("\n")
	var k = 0
	for {
		k++
		if k%2 == 0 {
			continue
		}
		fmt.Println("Ganjil doang nih =", k)

		if k >= 8 {
			break
		}
	}

	// Nested Loop
	fmt.Printf("\n")
	for l := 0; l < 10; l++ {
		for m := l; m < 10; m++ {
			fmt.Printf("%d", l)
		}
		fmt.Printf("\n")
	}

	// Labeling Loop
	// Can use break or continue when using labeling loop

lolping:
	for n := 0; n < 5; n++ {
		for o := 0; o < 5; o++ {
			if n == 4 {
				break lolping
			}
			fmt.Printf(`[ %d ][ %d ]`, n, o)
		}
		fmt.Printf("\n")
	}

	fmt.Println("\n============== Looping ==============")

}
