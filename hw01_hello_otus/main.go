package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	stringToReverse := "Hello, OTUS!"
	fmt.Println(stringutil.Reverse(stringToReverse))
}
