package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	f, _ := os.Open("test.al")
	defer f.Close()
	inp, _ := io.ReadAll(f)

	fmt.Println(string(inp))

	c := Compile(string(inp))
	fmt.Println(c.Objects)
}
