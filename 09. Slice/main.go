package main

import "fmt"

func main() {

	fmt.Println("============== Slice ==============")
	fmt.Printf("\n")
	// Slice is array reference
	// Using slice is like an array, but no need to write size
	// Example :
	var slice1 = []string{"Uzumaki", "Naruto"}
	fmt.Printf("%s\n", slice1)

	// Different between Array and Slice
	var thisSlice = []string{"Ini", "adalah", "SLICE"}                     // slice
	var thisArray1 = [2]string{"Ini", "Array"}                             // array
	var thisArray2 = [...]string{"Ini", "array", "tanpa", "nulis", "size"} // array
	fmt.Println(thisSlice, thisArray1, thisArray2)

	// How to call Slice and description
	var fruits = []string{"apple", "grape", "banana", "melon"}
	/*
		fruits[0:2]		[apple, grape]	semua elemen mulai indeks ke-0, hingga sebelum indeks ke-2
		fruits[0:4]		[apple, grape, banana, melon]	semua elemen mulai indeks ke-0, hingga sebelum indeks ke-4
		fruits[0:0]		[]	menghasilkan slice kosong, karena tidak ada elemen sebelum indeks ke-0
		fruits[4:4]		[]	menghasilkan slice kosong, karena tidak ada elemen yang dimulai dari indeks ke-4
		fruits[4:0]		[]	error, pada penulisan fruits[a,b] nilai a harus lebih besar atau sama dengan b
		fruits[:]		[apple, grape, banana, melon]	semua elemen
		fruits[2:]		[banana, melon]	semua elemen mulai indeks ke-2
		fruits[:2]		[apple, grape]	semua elemen hingga sebelum indeks ke-2
	*/
	fmt.Printf(`Ini isi semua slice bisa pakai "fruits[:]" atau "fruits" : %s%s`, fruits[:], "\n")
	fmt.Printf("%s\n", fruits[0:2])

	// Slice is reference type data
	// so when we we make another slice var with existing slice var
	// the new slice var is filled by existing slice var
	// which means each item of new slice are same memory location with the existing var
	// when we changing the item in exixting slice var, item on new slice is also changing
	// Example :

	var newFruits = fruits[1:3]
	fmt.Printf("New Fruit a => %s\n", newFruits)
	fmt.Println(cap(newFruits))

	var newFruitsb = newFruits[0:2]
	fmt.Printf("New Fruit b => %s\n", newFruitsb)
	fmt.Println(cap(newFruitsb))
	// cap for max length
	// len for length

	// append function
	var newSlice []string
	newSlice = append(newSlice, "Uzumaki")
	newSlice = append(newSlice, "Naruto")
	fmt.Printf("\n%s\n", newSlice)

	newSlice = append(newSlice, "Boruto")
	newSlice = append(newSlice, "Himawari")

	var newerSlice = newSlice[0:3]
	fmt.Printf("\n%d\n", cap(newerSlice))
	newerSlice = append(newerSlice, "Hyuga")
	fmt.Printf("\n%s\n", newerSlice)
	fmt.Printf("\n%s\n", newSlice)

	// copy function
	var variable1 = make([]string, 3)
	var variable2 = []string{"Uciha", "Sasuke"}
	_ = copy(variable1, variable2)
	variable1 = append(variable1, "Sarada")
	variable1 = append(variable1, "Sakura")

	fmt.Printf("Variable2 => %s\n", variable2)
	fmt.Printf("Variable1 => %s\n", variable1)
	fmt.Printf("Cap => %d\n", cap(variable1))

	for _, item := range fruits {
		fmt.Println(item)
	}

	fmt.Printf("\n")
	fmt.Println("============== Slice ==============")

}
