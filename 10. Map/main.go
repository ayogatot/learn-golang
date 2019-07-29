package main

import (
	"fmt"
)

func main() {
	fmt.Println("============== Map ==============")
	fmt.Printf("\n")

	// Map is like a Object on Javascript
	// must be initilaize
	// thisMap := map[string]int => error
	// thisMap := map[string]int{} => no error

	thisMonth := map[string]int{}
	thisMonth["januari"] = 1
	thisMonth["februari"] = 2
	fmt.Printf("Januari itu bulan ke - %d\n", thisMonth["januari"])

	arrMonth := [12]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	for i := 0; i < len(arrMonth); i++ {
		thisMonth[arrMonth[i]] = i + 1
	}
	fmt.Printf("%v \n", thisMonth)
	fmt.Printf("\n")

	for key, value := range thisMonth {
		fmt.Printf("%s: %d\n", key, value)
	}

	// delete map items
	fmt.Printf("\nJanuari sebelum di hapus => %d\n", thisMonth["januari"])
	delete(thisMonth, "januari")
	fmt.Printf("\nJanuari setelah di hapus => %d\n", thisMonth["januari"])
	// after delete value of key
	// zero value of int is 0

	// detect the key is exist or nay
	value, isExist := thisMonth["mei"]
	if isExist {
		fmt.Printf("%d\n", value)
	} else {
		fmt.Printf("Gada gan ! \n")
	}

	// Combination of map and slice is like an array of object in javascript
	// example :
	animals := []map[string]string{
		map[string]string{"jenis": "singo", "name": "michael"},
		map[string]string{"jenis": "iwak", "name": "iwak-ka"},
	}
	for _, animal := range animals {
		fmt.Println("Key :", animal["jenis"], "Value :", animal["name"])
	}

	fmt.Println("\n============== Map ==============")
}
