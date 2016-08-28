package main

import (
	"fmt"
	"os"
)

func main() {
	var r = 45
	fmt.Println("hello world")
	fmt.Fprintf(os.Stdout, "tttt %d", r)
}
