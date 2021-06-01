package main

import (
	"strings"
)

type Object struct {
	Def byte
	Pos Vec2
	Lasers []uint64
	Next []int // they are not represented as pointers, but as indexes in c.Objects
	Current bool
}

func (o *Object) IsMirror() bool {
	return strings.Contains("<>v^/\\", string(o.Def))
}
