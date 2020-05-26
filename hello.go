package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Printf("hello, my name is edgar\n")
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}