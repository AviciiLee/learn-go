package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(os.Args)
	os.Exit(-1)
}