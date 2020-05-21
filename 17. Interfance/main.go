package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diamter float64
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow((l.diamter/2), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diamter
}

func main() {
	fmt.Println("=====================Interface===================")
	var bangunDatar hitung

	bangunDatar = lingkaran{10.0}
	fmt.Println("=======Lingkaran")
	fmt.Println("Luas Lingkaran => ", bangunDatar.luas())
	fmt.Println("Keliling Lingkaran => ", bangunDatar.keliling())
}
