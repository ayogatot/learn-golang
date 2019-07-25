package main

import "fmt"

func main() {
	fmt.Println("============== IF & ELSE Condition ==============")
	// IF ELSE in go is same with other programming language
	// The difference is on bracket (parentless)
	// Go has no bracket () when we use IF ELSE
	const nilai = 80

	fmt.Printf("\nNilai kamu %d\n", nilai)
	if nilai > 95 && nilai <= 100 {
		fmt.Println("Lulus dengan nilai Sempurna !")
	} else if nilai > 75 && nilai <= 95 {
		fmt.Println("Lulus dengan nilai hampir Sempurna !")
	} else {
		fmt.Println("Selamat Anda tidak Lulus !")
	}

	// variable temporary at IF ELSE
	// can't use var
	// only interference is allowed to make temporary
	const point = 93.4

	fmt.Printf("\nPoint kamu sekarang %.4f\n", point)
	if precentage := point / 2; precentage > 50 {
		fmt.Printf("Poinmu Nice %f \n", precentage)
	} else if precentage > 30 && precentage < 50 {
		fmt.Printf("Pointmu nggak nice %f\n", precentage)
	}

	// condition using swith case :
	// break automatically even doesnt have break
	const number = 3
	switch number {
	case 1:
		fmt.Println("\nSelamat gan anda Juara 1")
	case 2:
		fmt.Println("\nAnda Juara 2 Ajg")
	default:
		fmt.Println("\nAnda itu Juara Harapan anjay")
	}

	// switch with IF ELSE style
	switch {
	case number > 1:
		{
			fmt.Println("YES")
		}
		// fallthrough
		// Can be use to force next case
	case number < 1:
		{
			fmt.Println("NO")
		}
	}

	fmt.Println("\n============== IF & ELSE Condition ==============")
}
