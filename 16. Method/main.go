package main

import "fmt"

type student struct {
	name  string
	grade int
}

/*
	func (s student) sayHello() maksudnya adalah fungsi sayHello dideklarasikan sebagai method milik struct student.
	Pada contoh di atas struct student memiliki dua buah method, yaitu sayHello() dan getGrade().
*/

func (s student) sayHello() {
	fmt.Println("halo", s.name)

}

func (s student) getGrade() {
	fmt.Println("Aku kelas", s.grade)
}

func (s *student) changeName(name string) {
	fmt.Println("change name to -->", name)
	s.name = name
}

func (s student) changeName2(name string) {
	fmt.Println("change name to -->", name)
	s.name = name
	// this method will not change the name in struct
	// it's just only can be use in this scope method
}

func main() {
	var s1 = student{"john wick", 21}
	s1.sayHello()
	s1.getGrade()
	s1.changeName("Jono")
	s1.sayHello()
}
