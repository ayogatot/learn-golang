package main

import "fmt"

type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("halo", s.name)

	/*
		func (s student) sayHello() maksudnya adalah fungsi sayHello dideklarasikan sebagai method milik struct student.
		Pada contoh di atas struct student memiliki dua buah method, yaitu sayHello() dan getNameAt().
	*/
}

func main() {
	var s1 = student{"john wick", 21}
	s1.sayHello()
}
