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

	c := compile(string(inp))

	c.initDefs()
	/*for _, o := range c.Objects {
		fmt.Println(&o)
	}*/

	for !c.ShouldQuit {
		c.tick()
	}
}

