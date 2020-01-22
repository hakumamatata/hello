package main

import (
	"fmt"
	"github.com/hakumamatata/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Printf("hello, world\n")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}