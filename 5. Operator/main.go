package main

import "fmt"

func main() {
	fmt.Println("OPERATOR")
	// ==================== A. Arithmetic Operator ====================
	/*
		+
		-
		/
		*
		%
	*/
	var result = ((12/3)*4 + 2) / 3
	fmt.Printf("This is result = %d \n", result)

	// ==================== B. Comparison Operator ====================
	/*
		==	apakah nilai kiri sama dengan nilai kanan
		!=	apakah nilai kiri tidak sama dengan nilai kanan
		<	apakah nilai kiri lebih kecil daripada nilai kanan
		<=	apakah nilai kiri lebih kecil atau sama dengan nilai kanan
		>	apakah nilai kiri lebih besar dari nilai kanan
		>=	apakah nilai kiri lebih besar atau sama dengan nilai kanan
	*/
	var value = result
	var equal = (value == 2)
	fmt.Printf("This operator = %t \n", equal)

	// ==================== B. Logic Operator ====================
	/*
		&&	kiri dan kanan
		||	kiri atau kanan
		!	negasi / nilai kebalikan
	*/

	const left = false
	const right = true

	const leftANDRigth = left && right
	fmt.Printf("Operator AND (&&) => %t\n", leftANDRigth)

	const leftORRigth = left || right
	fmt.Printf("Operator OR (||) => %t\n", leftORRigth)

	const notRigth = !right
	fmt.Printf("Operator Reverse (!) => %t\n", notRigth)

}
