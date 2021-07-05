package main

import (
	"os"
	"fmt"
	"strings"
)

type Compiler struct {
	Input []string
	Objects []Object
	Defs [256]Def
	CurrentTick uint64 // overflowing shouldn't break it, but just in case
	ShouldQuit bool
}

// Gets all object in the board
func (c *Compiler) getObjects() {
	for y := range c.Input {
		for x, s := range c.Input[y] {
			if s != ' ' {
				c.Objects = append(c.Objects, Object{Def: byte(s), Pos: Vec2{x, y}})
			}
		}
	}
}

// Checks, if v is in bounds
func (c *Compiler) inBounds(v Vec2) bool {
	return !(v.X < 0 || v.Y < 0 || v.Y >= len(c.Input) || v.X >= len(c.Input[v.Y]))
}

// Returns the byte on pos v
func (c *Compiler) onPos(v Vec2) byte {
	if !c.inBounds(v) {
		return ' '
	}

	return c.Input[v.Y][v.X]
}

// Returns an object on a pos v. O(n)
func (c *Compiler) objByPos(v Vec2) int {
	for i:=0; i < len(c.Objects); i++ {
		if c.Objects[i].Pos.Cmp(v) {
			return i
		}
	}
	return -1
}

// Walks in a direction v from start until it reaches the bounds or comes across an object
func (c *Compiler) walkDir(start Vec2, v Vec2) Vec2 {
	pos:=start.Add(v)
	for c.inBounds(pos) && c.onPos(pos) == ' ' { // this doesn't do the torus thingie
		pos = pos.Add(v)
	}
	return pos
}

// Generates objects from a graph
func (c *Compiler) genGraph() {
	for i := range c.Objects {
		c.Objects[i].Next = make([]int, 4)
		isMirror := c.Objects[i].isMirror()
		for j:=0; j < 4; j++ {
			c.Objects[i].Next[j] = -1

			next := c.walkDir(c.Objects[i].Pos, MOTIONS[j])
			if !isMirror {
				if !next.Cmp(c.Objects[i].Pos.Add(MOTIONS[j])) {
					continue
				}

				if c.onPos(next) != outMirrors[j] {
					continue
				}
			}

			c.Objects[i].Next[j] = c.objByPos(next)
		}
	}
}

func (c *Compiler) tick() {
	for i := range c.Objects {
		c.Objects[i].eval(c)
	}

	exit := true
	for i := range c.Objects {
		if len(c.Objects[i].Lasers) != 0 {
			exit = false
			break
		}
	}

	if exit {
		fmt.Printf("\nExiting because of no lasers.\nTotal ticks: %d.\n", c.CurrentTick + 1)
		os.Exit(0)
	}

	c.CurrentTick++
}

func compile(inp string) Compiler {
	c := Compiler{}
	c.Input = strings.Split(inp, "\n")
	c.getObjects()
	c.genGraph()
	c.ShouldQuit = false
	return c
}

