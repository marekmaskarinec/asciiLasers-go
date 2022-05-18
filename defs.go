package main

import (
	"fmt"
	"os"
)

type Def struct {
	Name string
	Char byte
	Argc int
	// returns a laser as the output, and true if it should be sent
	Func     func(o *Object, c *Compiler) (uint32, bool)
	WireFunc func(o *Object, c *Compiler)
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

func doLaserDetector(o *Object, c *Compiler) (uint32, bool) {
	o.Circuit.toggle(c)

	v := o.extractLasers(1, c.CurrentTick)
	if len(v) == 0 {
		return 0, false
	}

	return v[0], true
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

func wireNone(o *Object, c *Compiler) {}

func wireUpMirror(o *Object, c *Compiler) {
	if o.Circuit.Current {
		o.Def = 'v'
	} else {
		o.Def = '^'
	}
}

func wireRightMirror(o *Object, c *Compiler) {
	if o.Circuit.Current {
		o.Def = '<'
	} else {
		o.Def = '>'
	}
}

func wireLeftMirror(o *Object, c *Compiler) {
	if o.Circuit.Current {
		o.Def = '>'
	} else {
		o.Def = '<'
	}
}

func wireDownMirror(o *Object, c *Compiler) {
	if o.Circuit.Current {
		o.Def = '^'
	} else {
		o.Def = 'v'
	}
}

func wireDeleter(o *Object, c *Compiler) {
	if o.Circuit.Current {
		o.Def = '*'
	} else {
		o.Def = '#'
	}
}

func wireReflector(o *Object, c *Compiler) {
	if o.Circuit.Current {
		o.Def = '#'
	} else {
		o.Def = '*'
	}
}

func wireIncrement(o *Object, c *Compiler) {
	if o.Circuit.Current {
		o.Def = 'd'
	} else {
		o.Def = 'i'
	}
}

func wireDecrement(o *Object, c *Compiler) {
	if o.Circuit.Current {
		o.Def = 'i'
	} else {
		o.Def = 'd'
	}
}

func wireAddition(o *Object, c *Compiler) {
	if o.Circuit.Current {
		o.Def = 's'
	} else {
		o.Def = 'a'
	}
}

func wireSubtraction(o *Object, c *Compiler) {
	if o.Circuit.Current {
		o.Def = 'a'
	} else {
		o.Def = 's'
	}
}

func wireMultiplication(o *Object, c *Compiler) {
	if o.Circuit.Current {
		o.Def = 'n'
	} else {
		o.Def = 'm'
	}
}

func wireDivision(o *Object, c *Compiler) {
	if o.Circuit.Current {
		o.Def = 'm'
	} else {
		o.Def = 'n'
	}
}

func (c *Compiler) initDefs() {
	// IO
	c.Defs['$'] = Def{"value print", '$', 1, doValuePrint, nil}
	c.Defs['&'] = Def{"char print", '&', 1, doCharPrint, nil}
	c.Defs['{'] = Def{"program start", '{', 0, doProgramStart, nil}
	c.Defs['}'] = Def{"program end", '}', -1, doProgramEnd, nil}
	c.Defs['_'] = Def{"input", '_', -1, doInput, nil}

	// mirrors
	c.Defs['\\'] = Def{"backslash mirror", '\\', -1, nil, nil}
	c.Defs['/'] = Def{"slash mirror", '/', -1, nil, nil}
	c.Defs['^'] = Def{"up mirror", '^', -1, doUpMirror, wireUpMirror}
	c.Defs['>'] = Def{"right mirror", '>', -1, doRightMirror, wireRightMirror}
	c.Defs['v'] = Def{"down mirror", 'v', -1, doDownMirror, wireDownMirror}
	c.Defs['<'] = Def{"left mirror", '<', -1, doLeftMirror, wireLeftMirror}
	c.Defs['='] = Def{"horizontal mirror", '=', -1, nil, nil}
	c.Defs['H'] = Def{"vertical mirror", 'H', -1, nil, nil}

	// modifiers
	c.Defs['*'] = Def{"reflector", '*', 1, doReflector, wireReflector}
	c.Defs['i'] = Def{"increment", 'i', 1, doIncrement, wireIncrement}
	c.Defs['d'] = Def{"decrement", 'd', 1, doDecrement, wireDecrement}
	c.Defs['@'] = Def{"laser detector", '@', 1, doLaserDetector, nil}

	c.Defs['0'] = Def{"set to 0", '0', 1, do0, nil}
	c.Defs['1'] = Def{"set to 1", '1', 1, do1, nil}
	c.Defs['2'] = Def{"set to 2", '2', 1, do2, nil}
	c.Defs['3'] = Def{"set to 3", '3', 1, do3, nil}
	c.Defs['4'] = Def{"set to 4", '4', 1, do4, nil}
	c.Defs['5'] = Def{"set to 5", '5', 1, do5, nil}
	c.Defs['6'] = Def{"set to 6", '6', 1, do6, nil}
	c.Defs['7'] = Def{"set to 7", '7', 1, do7, nil}
	c.Defs['8'] = Def{"set to 8", '8', 1, do8, nil}
	c.Defs['9'] = Def{"set to 9", '9', 1, do9, nil}
	c.Defs['A'] = Def{"set to 10", 'A', 1, doA, nil}
	c.Defs['B'] = Def{"set to 11", 'B', 1, doB, nil}
	c.Defs['C'] = Def{"set to 12", 'C', 1, doC, nil}
	c.Defs['D'] = Def{"set to 13", 'D', 1, doD, nil}
	c.Defs['E'] = Def{"set to 14", 'E', 1, doE, nil}
	c.Defs['F'] = Def{"set to 15", 'F', 1, doF, nil}

	c.Defs['#'] = Def{"deleter", '#', 1, doDeleter, wireDeleter}
	c.Defs['m'] = Def{"multiplication", 'm', 2, doMultiplication, wireMultiplication}
	c.Defs['n'] = Def{"division", 'n', 2, doDivision, wireDivision}
	c.Defs['a'] = Def{"addition", 'a', 2, doAddition, wireAddition}
	c.Defs['s'] = Def{"subtraction", 's', 2, doSubtraction, wireSubtraction}
	c.Defs['l'] = Def{"modulo", 'l', 2, doModulo, nil}
}
