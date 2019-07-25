package main

import "fmt"

func main() {
	// data type in go : numeric (decimal or non-decimal), string, boolean
	// ================ A. Numeric Non-Decimal or Non Floating Point =====================
	// uint = positive numeric
	// int = negative and positive numeric
	/*
		uint8	0 ↔ 255
		uint16	0 ↔ 65535
		uint32	0 ↔ 4294967295
		uint64	0 ↔ 18446744073709551615
		uint	sama dengan uint32 atau uint64 (tergantung nilai)
		byte	sama dengan uint8
		int8	-128 ↔ 127
		int16	-32768 ↔ 32767
		int32	-2147483648 ↔ 2147483647
		int64	-9223372036854775808 ↔ 9223372036854775807
		int	sama dengan int32 atau int64 (tergantung nilai)
		rune	sama dengan int32
	*/
	// Use type wisely, because can effect allocation of memory
	var positiveNumber uint8 = 200
	fmt.Printf("This is positive number = %d \n", positiveNumber)
	// %d is a template to place or to format numeric data non-decimal

	// ================ B. Numeric Decimal or Floating Point =====================
	var decimalNumber = 2.52

	fmt.Printf("This is decimal = %f\n", decimalNumber)
	fmt.Printf("This is decimal = %.3f\n", decimalNumber)
	fmt.Printf("This is decimal = %.10f\n", decimalNumber)
	// %f is a template to format numeric decimal to string
	// normal digit is 6
	// %.nf "n" can be replace number that represent the digit number

	// ================ C. Boolean =====================
	var thisBool bool
	thisBool = true

	fmt.Printf("This is boolean = %t \n", thisBool)
	// %t is a template to format boolean to string

	// ================ D. String =====================
	var message string
	message = "Hello"

	fmt.Printf("This is string = %s\n", message)
	// %s is a template to place variable string

	message = `Hello myName  is "John Wick".
	My Age now is 22
	`
	fmt.Println(message)

	// ================ E. Zero Value or Nil ================
	// nil is same with null , empty or zero value
	// variable with type data which can be set with nil is :
	/*
		pointer
		tipe data fungsi
		slice
		map
		channel
		interface kosong atau interface{}
	*/
	// every data type have zero value
	/*
		Zero value dari string adalah "" (string kosong).
		Zero value dari bool adalah false.
		Zero value dari tipe numerik non-desimal adalah 0.
		Zero value dari tipe numerik desimal adalah 0.0.
	*/

}
