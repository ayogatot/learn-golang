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
	var s1 student
	s1.name = "john wick"
	s1.grade = 2

	fmt.Println("name  :", s1.name)
	fmt.Println("grade :", s1.grade)
}
