package main

import (
	"fmt"
	"strings"
)

type Laser struct {
	Val  uint32
	Tick uint64
}

type Circuit struct {
	// indices to the object array
	Objects []int
	Current bool
}

func (c *Circuit) toggle(com *Compiler) {
	c.Current = !c.Current

	for i := 0; i < len(c.Objects); i++ {
		fn := com.Defs[com.Objects[c.Objects[i]].Def].WireFunc
		if fn != nil {
			fn(&com.Objects[c.Objects[i]], com)
		}
	}
}

type Object struct {
	Def    byte
	Pos    Vec2
	Lasers []Laser
	// indices to the object array
	Next    [4]int
	Circuit *Circuit
}

func (o *Object) String() string {
	return fmt.Sprintf("[ %d %d ]: { '%c', %v, %d }", o.Pos.X, o.Pos.Y, o.Def, o.Next, len(o.Circuit.Objects))
}

func (o *Object) isMirror() bool {
	return strings.Contains("<>v^", string(o.Def))
}

func (o *Object) isWire() bool {
	return strings.Contains("-|+O", string(o.Def))
}

func (o *Object) extractLasers(count int, tick uint64) []uint32 {
	out := []uint32{}

	for i := 0; (i < count || count == -1) && i < len(o.Lasers) && o.Lasers[i].Tick != tick; i++ {
		out = append(out, o.Lasers[i].Val)
	}

	if count == -1 {
		o.Lasers = []Laser{}
	}

	if len(o.Lasers) > 0 && (len(out) == count || count == -1) {
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
