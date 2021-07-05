package main

import (
	"fmt"
	"os"
)

type Def struct {
	Name string
	Char byte
	Argc int // 0 means that it is operated by current. Negative values mean, it can be operated by both + argc
	Func func(o *Object, c *Compiler) (uint32, bool) // returns a laser as the output, and true if it should be sent
}

// IO
func doValuePrint(o *Object, c *Compiler) (uint32, bool) {
	lasers := o.extractLasers(-1, c.CurrentTick)
	for _, v := range lasers {
		fmt.Println(v)
	}

	return 0, false
}

func doProgramStart(o *Object, c *Compiler) (uint32, bool) {
	if c.CurrentTick == 0 {
		return 1, true
	}
	return 0, false
}

func doProgramEnd(o *Object, c *Compiler) (uint32, bool) {
	v := o.extractLasers(1, c.CurrentTick)
	if len(v) == 0 {
		os.Exit(-1)
	}
	os.Exit(int(v[0]))

	return 0, false
}

// modifiers
func doReflector(o *Object, c *Compiler) (uint32, bool) {
	v := o.extractLasers(1, c.CurrentTick)
	if len(v) == 0 {
		return 0, false
	}

	return v[0], true
}

//mirrors
func doMirror(o *Object, c *Compiler, dir int) {
	v := o.extractLasers(1, c.CurrentTick)
	if len(v) == 1 && o.Next[dir] >= 0 {
		c.Objects[o.Next[dir]].Lasers = append(c.Objects[o.Next[dir]].Lasers, Laser{v[0], c.CurrentTick})
	}
}

func doUpMirror(o *Object, c *Compiler) (uint32, bool) {
	doMirror(o, c, UP)
	return 0, false
}

func doRightMirror(o *Object, c *Compiler) (uint32, bool) {
	doMirror(o, c, RIGHT)
	return 0, false
}

func doDownMirror(o *Object, c *Compiler) (uint32, bool) {
	doMirror(o, c, DOWN)
	return 0, false
}

func doLeftMirror(o *Object, c *Compiler) (uint32, bool) {
	doMirror(o, c, LEFT)
	return 0, false
}

func (c *Compiler) initDefs() {
	// IO
	c.Defs['$'] = Def{"value print", '$', 1, doValuePrint}
	c.Defs['&'] = Def{"char print", '&', 1, nil}
	c.Defs['{'] = Def{"program start", '{', 0, doProgramStart}
	c.Defs['}'] = Def{"program end", '}', -1, doProgramEnd}
	c.Defs['_'] = Def{"input", '_', -1, nil}

	// mirrors
	c.Defs['\\'] = Def{"backslash mirror", '\\', -1, nil}
	c.Defs['/'] = Def{"slash mirror", '/', -1, nil}
	c.Defs['^'] = Def{"up mirror", '^', -1, doUpMirror}
	c.Defs['>'] = Def{"right mirror", '>', -1, doRightMirror}
	c.Defs['v'] = Def{"down mirror", 'v', -1, doDownMirror}
	c.Defs['<'] = Def{"left mirror", '<', -1, doLeftMirror}
	c.Defs['='] = Def{"horizontal mirror", '=', -1, nil}
	c.Defs['H'] = Def{"vertical mirror", 'H', -1, nil}
	
	// modifiers
	c.Defs['*'] = Def{"reflector", '*', 1, doReflector}
	c.Defs['i'] = Def{"increment", 'i', 1, nil}
	c.Defs['d'] = Def{"decrement", 'd', 1, nil}
	// 0 - F is handled as a special case
	c.Defs['#'] = Def{"deleter", '#', 1, nil}
	c.Defs['m'] = Def{"multiplication", 'm', 2, nil}
	c.Defs['n'] = Def{"division", 'n', 2, nil}
	c.Defs['a'] = Def{"addition", 'a', 2, nil}
	c.Defs['s'] = Def{"subtraction", 's', 2, nil}
	c.Defs['l'] = Def{"modulo", 'l', 2, nil}

	// wires are TODO
}
