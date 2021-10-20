package main

import "fmt"

const HEIGHT = 200

func main() {
	const (
		C1 = "C1C1C1"
		C2 = "C2C2C2"
		C3 = "C3C3C3"
	)

	fmt.Println(HEIGHT)
	fmt.Println(C1)

	s1 := "My String"
	var s2 = "My String"
	var s3 string = "My String"

	fmt.Println(s1, s2, s3)

	const s4 = "My String"
	const s5 string = "My String"

	fmt.Println(s4, s5)

	const s6 = 123
	const s7 float64 = 123
	fmt.Println(s6, s7)

	// var v1 float32 = s6 * 12
	// var v2 float32 = s7 * 12
}
