package main

type Def struct {
	Name string
	Char byte
	Argc int // 0 means that it is operated by current. Negative values mean, it can be operated by both + argc
	Func func(o *Object) (uint64, int) // laser return value and direction. if direction is -1, it means laser has been fired, or there is no need to return one.
}

func (c *Compiler) InitDefs() { // this is gonna be painful
	// IO
	c.Defs['$'] = Def{"value print", '$', 1, nil}
	c.Defs['&'] = Def{"char print", '&', 1, nil}
	c.Defs['{'] = Def{"program start", '{', 0, nil}
	c.Defs['}'] = Def{"program end", '}', -1, nil}
	c.Defs['_'] = Def{"input", '_', -1, nil}

	// mirrors
	// they should probably be a special case
	c.Defs['\\'] = Def{"backslash mirror", '\\', -1, nil}
	c.Defs['/'] = Def{"slash mirror", '/', -1, nil}
	c.Defs['^'] = Def{"up mirror", '^', -1, nil}
	c.Defs['>'] = Def{"right mirror", '>', -1, nil}
	c.Defs['v'] = Def{"down mirror", 'v', -1, nil}
	c.Defs['<'] = Def{"left mirror", '<', -1, nil}
	c.Defs['='] = Def{"horizontal mirror", '=', -1, nil}
	c.Defs['H'] = Def{"vertical mirror", 'H', -1, nil}
	
	// modifiers
	c.Defs['*'] = Def{"reflector", '*', 1, nil} // why is this a modifier?
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
