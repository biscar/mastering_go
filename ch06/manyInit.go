package main

import (
	"b"
	"a"
	"fmt"
)

func init() {
	fmt.Println("init() manyInit")
}

func main() {
	b.FromB()
	a.FromA()
}
