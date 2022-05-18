package main

import (
	"fmt"
	"strings"
	"unicode"
)

func printWarning(message string, o Object) {
	fmt.Printf("\033[35m\033[1mwarning:\033[0m Object %c at %d %d %s.\n", o.Def, o.Pos.X, o.Pos.Y, message)
}

func (c *Compiler) warn() {
	c.checkForUselessBlocks()
	c.checkForInvisibles()
}

func (c *Compiler) checkForUselessBlocks() {
	for _, o := range c.Objects {
		if strings.Contains("$&#", string(o.Def)) || o.isWire() {
			continue
		}

		isOk := false
		for _, n := range o.Next {
			if n != -1 {
				isOk = true
				break
			}

		}
		if !isOk {
			printWarning("has no outputs", o)
		}
	}
}

func (c *Compiler) checkForInvisibles() {
	for y := range c.Input {
		for x, b := range c.Input[y] {
			if unicode.IsControl(b) {
				printWarning("is a control character", Object{Def: byte(b), Pos: Vec2{x, y}})
			}
		}
	}
}
