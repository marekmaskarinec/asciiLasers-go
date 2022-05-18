package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	f, _ := os.Open("test.al")
	defer f.Close()
	inp, _ := io.ReadAll(f)

	fmt.Println(string(inp))

	c := compile(string(inp))

	c.initDefs()
	for i, o := range c.Objects {
		fmt.Println(i, &o)
	}

	startTime = time.Now()
	for !c.ShouldQuit {
		c.prettyPrint()
		time.Sleep(200 * time.Millisecond)
		fmt.Print("\033[1J")
		c.tick()
	}
}
