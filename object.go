package main

import (
	"strings"
	"fmt"
)

type Laser struct {
	Val uint32
	Tick uint64
}

type Object struct {
	Def byte
	Pos Vec2
	Lasers []Laser
	Next []int // they are not represented as pointers, but as indexes in c.Objects
	Current bool
}

func (o *Object) String() string {
	return fmt.Sprintf("[ %d %d ]: { '%c', %v }", o.Pos.X, o.Pos.Y, o.Def, o.Next)
}

func (o *Object) isMirror() bool {
	return strings.Contains("<>v^/\\", string(o.Def))
}

func (o *Object) extractLasers(count int, tick uint64) []uint32 {
	out := []uint32{}

	for i:=0; (i < count || count == -1) && i < len(o.Lasers); i++ {
		out = append(out, o.Lasers[i].Val)
	}

	if len(out) == count {
		o.Lasers = o.Lasers[len(out):]
	}
	return out
}

func (o *Object) eval(c *Compiler) {
	if c.Defs[o.Def].Func == nil {
		fmt.Printf("%c not implemented\n", o.Def)
		return
	}

	//fmt.Println(o)
	//fmt.Println(o.Lasers)

	if outl, shouldSend := c.Defs[o.Def].Func(o, c); shouldSend {
		for _, n := range o.Next {
			if n < 0 {
				continue
			}
			c.Objects[n].Lasers = append(c.Objects[n].Lasers, Laser{outl, c.CurrentTick})
		}
	}
}
