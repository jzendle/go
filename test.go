package main

import (
	"fmt"
	"os"
)

func main() {
	var r = 99
	fmt.Println("hello world")
	fmt.Fprintf(os.Stdout, "tttt %d", r)
}
