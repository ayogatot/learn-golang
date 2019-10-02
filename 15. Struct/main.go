package main

import "fmt"

type student struct {
	name  string
	grade int
}

/*
Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method),
yang dibungkus sebagai tipe data baru dengan nama tertentu. Property dalam struct,
tipe datanya bisa bervariasi. Mirip seperti map, hanya saja key-nya sudah didefinisikan di awal,
dan tipe data tiap itemnya bisa berbeda.
*/

/*
Dari sebuah struct, kita bisa buat variabel baru, yang memiliki atribut sesuai skema struct tersebut.
Kita sepakati dalam buku ini, variabel tersebut dipanggil dengan istilah object atau object struct.
*/

func main() {
	var s1 = student{}
	s1.name = "wick"
	s1.grade = 2

	var s2 = student{"ethan", 2}

	var s3 = student{name: "jason"}

	fmt.Println("student 1 :", s1.name)
	fmt.Println("student 2 :", s2.name)
	fmt.Println("student 3 :", s3.name)

	/*
		1. Deklarasi s1 = manampung object student di variable s1, dan menambahkan property nya dibawahnya.
		2. Deklarasi s2 = sama dengan s1, namun langsung mengisi property-nya pada saat pembuatan variable
		2. Deklarasi s3 = sama dengan s2, namun hanya mengisi 1 property saya, cara ini efektif ketika membuat struct namun tidak semua property harus memiliki nilai
	*/

	var s4 = student{name: "wayne", grade: 2}
	var s5 = student{grade: 2, name: "bruce"}

	// Kita juga bisa mendeklarasi struct dengan mengisi property-nya secara tidak urut

	fmt.Println("student 1 :", s4.name)
	fmt.Println("student 1 :", s5.name)

	// Variable object Pointer
	var s6 *student = &s1

	fmt.Println("student 1 :", s6.name)

}
