package main

import "fmt"

func main() {
	fmt.Println("ku")

	// zeroNum := 0
	zeroFNum := 0.0

	fnum1 := 8 / zeroNum 
	fmt.Println(fnum1) // panic: runtime error: integer divide by zero

	fnum2 := 8 / zeroFNum
	fmt.Println(fnum2) // invalid operation: division by zero
}
