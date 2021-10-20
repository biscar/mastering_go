package main

import "fmt"

func main() {
	aMap := map[string]int{}
	// var aMap map[string]int
	aMap = nil
	fmt.Println(aMap)
	aMap["test"] = 1
}
