package main

import (
	"fmt"
	"os"
)

type Def struct {
	Name string
	Char byte
	Argc int                                         // 0 means that it is operated by current. Negative values mean, it can be operated by both + argc
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

func doCharPrint(o *Object, c *Compiler) (uint32, bool) {
	lasers := o.extractLasers(-1, c.CurrentTick)
	for _, v := range lasers {
		fmt.Printf("%c", v)
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

func doInput(o *Object, c *Compiler) (uint32, bool) {
	var out uint32
	fmt.Scanf("%u", &out)
	return out, true
}

// modifiers
func doReflector(o *Object, c *Compiler) (uint32, bool) {
	v := o.extractLasers(1, c.CurrentTick)
	if len(v) == 0 {
		return 0, false
	}

	return v[0], true
}

func doIncrement(o *Object, c *Compiler) (uint32, bool) {
	v := o.extractLasers(1, c.CurrentTick)
	if len(v) == 0 {
		return 0, false
	}

	return v[0] + 1, true
}

func doDecrement(o *Object, c *Compiler) (uint32, bool) {
	v := o.extractLasers(1, c.CurrentTick)
	if len(v) == 0 {
		return 0, false
	}

	return v[0] - 1, true
}

func doDeleter(o *Object, c *Compiler) (uint32, bool) {
	return 0, false
}

func doSetValue(o *Object, c *Compiler, val uint32) (uint32, bool) {
	v := o.extractLasers(1, c.CurrentTick)
	if len(v) == 0 {
		return 0, false
	}
	return val, true
}

func do0(o *Object, c *Compiler) (uint32, bool) {
	return doSetValue(o, c, 0)
}

func do1(o *Object, c *Compiler) (uint32, bool) {
	return doSetValue(o, c, 1)
}

func do2(o *Object, c *Compiler) (uint32, bool) {
	return doSetValue(o, c, 2)
}

func do3(o *Object, c *Compiler) (uint32, bool) {
	return doSetValue(o, c, 3)
}

func do4(o *Object, c *Compiler) (uint32, bool) {
	return doSetValue(o, c, 4)
}

func do5(o *Object, c *Compiler) (uint32, bool) {
	return doSetValue(o, c, 5)
}

func do6(o *Object, c *Compiler) (uint32, bool) {
	return doSetValue(o, c, 6)
}

func do7(o *Object, c *Compiler) (uint32, bool) {
	return doSetValue(o, c, 7)
}

func do8(o *Object, c *Compiler) (uint32, bool) {
	return doSetValue(o, c, 8)
}

func do9(o *Object, c *Compiler) (uint32, bool) {
	return doSetValue(o, c, 9)
}

func doA(o *Object, c *Compiler) (uint32, bool) {
	return doSetValue(o, c, 10)
}

func doB(o *Object, c *Compiler) (uint32, bool) {
	return doSetValue(o, c, 11)
}

func doC(o *Object, c *Compiler) (uint32, bool) {
	return doSetValue(o, c, 12)
}

func doD(o *Object, c *Compiler) (uint32, bool) {
	return doSetValue(o, c, 13)
}

func doE(o *Object, c *Compiler) (uint32, bool) {
	return doSetValue(o, c, 14)
}

func doF(o *Object, c *Compiler) (uint32, bool) {
	return doSetValue(o, c, 14)
}

func doMath(o *Object, c *Compiler, m func(a, b uint32) uint32) (uint32, bool) {
	v := o.extractLasers(2, c.CurrentTick)
	if len(v) != 2 {
		return 0, false
	}

	return m(v[0], v[1]), true
}

func doMultiplication(o *Object, c *Compiler) (uint32, bool) {
	return doMath(o, c, func(a, b uint32) uint32 { return a * b })
}

func doDivision(o *Object, c *Compiler) (uint32, bool) {
	return doMath(o, c, func(a, b uint32) uint32 { return a / b })
}

func doAddition(o *Object, c *Compiler) (uint32, bool) {
	return doMath(o, c, func(a, b uint32) uint32 { return a + b })
}

func doSubtraction(o *Object, c *Compiler) (uint32, bool) {
	return doMath(o, c, func(a, b uint32) uint32 { return a - b })
}

func doModulo(o *Object, c *Compiler) (uint32, bool) {
	return doMath(o, c, func(a, b uint32) uint32 { return a % b })
}

//mirrors
func doMirror(o *Object, c *Compiler, dir int) {
	v := o.extractLasers(1, c.CurrentTick)
	if len(v) == 1 && o.Next[dir] >= 0 {
		c.Objects[o.Next[dir]].Lasers = append(c.Objects[o.Next[dir]].Lasers, Laser{v[0], c.CurrentTick + 1})
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
	c.Defs['&'] = Def{"char print", '&', 1, doCharPrint}
	c.Defs['{'] = Def{"program start", '{', 0, doProgramStart}
	c.Defs['}'] = Def{"program end", '}', -1, doProgramEnd}
	c.Defs['_'] = Def{"input", '_', -1, doInput}

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
	c.Defs['i'] = Def{"increment", 'i', 1, doIncrement}
	c.Defs['d'] = Def{"decrement", 'd', 1, doDecrement}

	c.Defs['0'] = Def{"set to 0", '0', 1, do0}
	c.Defs['1'] = Def{"set to 1", '1', 1, do1}
	c.Defs['2'] = Def{"set to 2", '2', 1, do2}
	c.Defs['3'] = Def{"set to 3", '3', 1, do3}
	c.Defs['4'] = Def{"set to 4", '4', 1, do4}
	c.Defs['5'] = Def{"set to 5", '5', 1, do5}
	c.Defs['6'] = Def{"set to 6", '6', 1, do6}
	c.Defs['7'] = Def{"set to 7", '7', 1, do7}
	c.Defs['8'] = Def{"set to 8", '8', 1, do8}
	c.Defs['9'] = Def{"set to 9", '9', 1, do9}
	c.Defs['A'] = Def{"set to 10", 'A', 1, doA}
	c.Defs['B'] = Def{"set to 11", 'B', 1, doB}
	c.Defs['C'] = Def{"set to 12", 'C', 1, doC}
	c.Defs['D'] = Def{"set to 13", 'D', 1, doD}
	c.Defs['E'] = Def{"set to 14", 'E', 1, doE}
	c.Defs['F'] = Def{"set to 15", 'F', 1, doF}

	c.Defs['#'] = Def{"deleter", '#', 1, doDeleter}
	c.Defs['m'] = Def{"multiplication", 'm', 2, doMultiplication}
	c.Defs['n'] = Def{"division", 'n', 2, doDivision}
	c.Defs['a'] = Def{"addition", 'a', 2, doAddition}
	c.Defs['s'] = Def{"subtraction", 's', 2, doSubtraction}
	c.Defs['l'] = Def{"modulo", 'l', 2, doModulo}

	// wires are TODO
}
