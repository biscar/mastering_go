package main

import (
	"fmt"
)

func main() {
	aSliceLiteral := []int{1, 2, 3, 4, 5}
	integer := make([]int, 20)
	fmt.Println(cap(aSliceLiteral))

	fmt.Printf("type %T\n", aSliceLiteral)
	fmt.Printf("type %T\n", integer)
	fmt.Println(aSliceLiteral[0], "", integer[10])

	for i := 0; i < len(integer); i++ {
		fmt.Println(integer[i])
	}

	aSliceLiteral = nil

	for i := 0; i < len(aSliceLiteral); i++ {
		fmt.Println(aSliceLiteral[i])
	}
	fmt.Printf("type %T\n", aSliceLiteral)

	integer = append(integer, 12345)
	fmt.Println(integer[20])

	fmt.Println(integer[19:21])
	fmt.Println("```````````")

	s1 := make([]int, 5)
	reSlice := s1[1:3]
	fmt.Println(s1)
	fmt.Println(reSlice)
	reSlice[0] = -100
	reSlice[1] = 123456
	fmt.Println(s1)
	fmt.Println(reSlice)

}
