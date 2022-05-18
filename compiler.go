package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var startTime time.Time

type Compiler struct {
	Input       []string
	Objects     []Object
	Defs        [256]Def
	CurrentTick uint64 // overflowing shouldn't break it, but just in case
	ShouldQuit  bool
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
	for i := 0; i < len(c.Objects); i++ {
		if c.Objects[i].Pos.Cmp(v) {
			return i
		}
	}
	return -1
}

// Walks in a direction v from start until it reaches the bounds or comes across an object
func (c *Compiler) walkDir(start Vec2, v Vec2) Vec2 {
	pos := start.Add(v)
	for c.inBounds(pos) && (c.onPos(pos) == ' ' || isWire(c.onPos(pos))) {
		pos = pos.Add(v)
	}
	return pos
}

func (c *Compiler) walkWireDir(p1, v Vec2) Vec2 {
	p2 := p1.Add(v)
	c1 := c.onPos(p1)
	c2 := c.onPos(p2)

	if (p1.X-p2.X != 0 && (c1 == '-' || c1 == 'O') && (c2 == '-' || c2 == 'O')) ||
		(p1.Y-p2.Y != 0 && (c1 == '|' || c1 == 'O') && (c2 == '|' || c2 == 'O')) ||
		(!isWire(c1) && isWire(c2)) || (isWire(c1) && !isWire(c2)) {
		return p2
	}

	if c2 == '+' {
		return c.walkWireDir(p1, v.Add(v))
	}

	return Vec2{-1, -1}
}

func (c *Compiler) genCircuit(idx int, wireGraph [][4]int) {
	o := &c.Objects[idx]

	if o.Circuit == nil {
		o.Circuit = new(Circuit)
		o.Circuit.Objects = []int{idx}
	}

	for i := 0; i < 4; i++ {
		if wireGraph[idx][i] < 0 {
			continue
		}

		if c.Objects[wireGraph[idx][i]].Circuit != nil {
			continue
		}

		o.Circuit.Objects = append(o.Circuit.Objects, wireGraph[idx][i])
		c.Objects[wireGraph[idx][i]].Circuit = o.Circuit
		c.genCircuit(wireGraph[idx][i], wireGraph)
	}
}

// Generates objects from a graph
func (c *Compiler) genGraph() {
	wireGraph := make([][4]int, len(c.Objects))

	for i := range c.Objects {
		isMirror := c.Objects[i].isMirror()
		for j := 0; j < 4; j++ {
			c.Objects[i].Next[j] = -1
			wireGraph[i][j] = -1

			wireGraph[i][j] = c.objByPos(
				c.walkWireDir(c.Objects[i].Pos, MOTIONS[j]))

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

	for i := 0; i < len(wireGraph); i++ {
		if c.Objects[i].Circuit != nil {
			continue
		}

		c.genCircuit(i, wireGraph)
	}
}

func (c *Compiler) prettyPrint() {
	for y := 0; y < len(c.Input); y++ {
		for x := 0; x < len(c.Input[y]); x++ {
			if c.Input[y][x] == ' ' {
				fmt.Print(" ")
				continue
			}

			o := c.Objects[c.objByPos(Vec2{x, y})]
			if len(o.Lasers) != 0 {
				fmt.Print("\033[41m")
			}
			if o.Circuit.Current {
				fmt.Print("\033[33m")
			}
			fmt.Printf("%c\033[0m", o.Def)
		}
		fmt.Println()
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

	if c.CurrentTick > 100000 || exit {
		now := time.Now()
		fmt.Printf("\nExiting because of no lasers.\nTotal ticks: %d.\nTicks per second: %f.\n", c.CurrentTick+1, float32(c.CurrentTick+1)/float32(now.Sub(startTime).Nanoseconds())*1000000000)
		os.Exit(0)
	}

	c.CurrentTick++
}

func compile(inp string) Compiler {
	c := Compiler{}
	c.Input = strings.Split(inp, "\n")
	c.getObjects()
	c.genGraph()
	c.warn()
	c.ShouldQuit = false
	return c
}
