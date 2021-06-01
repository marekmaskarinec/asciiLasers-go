package main

import (
	"strings"
)

type Compiler struct {
	Input []string
	Objects []Object
	Defs [256]Def
}

// Gets all object in the board
func (c *Compiler) GetObjects() {
	for y := range c.Input {
		for x, s := range c.Input[y] {
			if s != ' ' {
				c.Objects = append(c.Objects, Object{Def: byte(s), Pos: Vec2{x, y}})
			}
		}
	}
}

// Checks, if v is in bounds
func (c *Compiler) InBounds(v Vec2) bool {
	return !(v.X < 0 || v.Y < 0 || v.Y >= len(c.Input) || v.X >= len(c.Input[v.Y]))
}

// Returns the byte on pos v
func (c *Compiler) OnPos(v Vec2) byte {
	if !c.InBounds(v) {
		return ' '
	}

	return c.Input[v.Y][v.X]
}

// Returns an object on a pos v. O(n)
func (c *Compiler) ObjByPos(v Vec2) int {
	for i:=0; i < len(c.Objects); i++ {
		if c.Objects[i].Pos.Cmp(v) {
			return i
		}
	}
	return -1
}

// Walks in a direction v from start until it reaches the bounds or comes across an object
func (c *Compiler) WalkDir(start Vec2, v Vec2) Vec2 {
	pos:=start.Add(v)
	for c.InBounds(pos) && c.OnPos(pos) == ' ' { // this doesn't do the torus thingie
		pos = pos.Add(v)
	}
	return pos
}

// Generates objects from a graph
func (c *Compiler) GenGraph() {
	for i := range c.Objects {
		c.Objects[i].Next = make([]int, 4)
		isMirror := c.Objects[i].IsMirror()
		for j:=0; j < 4; j++ {
			c.Objects[i].Next[j] = -1

			next := c.WalkDir(c.Objects[i].Pos, MOTIONS[j])
			if !isMirror {
				if !next.Cmp(c.Objects[i].Pos.Add(MOTIONS[j])) {
					continue
				}

				if c.OnPos(next) != outMirrors[j] {
					continue
				}
			}

			c.Objects[i].Next[j] = c.ObjByPos(next)
		}
	}
}

func Compile(inp string) Compiler {
	c := Compiler{}
	c.Input = strings.Split(inp, "\n")
	c.GetObjects()
	c.GenGraph()
	return c
}
