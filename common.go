package main

//x---'
//    'Y
type Vec2 struct {
	X int
	Y int
}

func (v *Vec2) Add(v2 Vec2) Vec2 { // this doesn't modify
	return Vec2{v.X + v2.X, v.Y + v2.Y}
}

func (v *Vec2) Cmp(v2 Vec2) bool {
	return (v.X == v2.X && v.Y == v2.Y)
}

const UP = 0
const DOWN = 1
const LEFT = 2
const RIGHT = 3

var VUP = Vec2{X: 0, Y: -1}
var VDOWN = Vec2{X: 0, Y: 1}
var VLEFT = Vec2{X: -1, Y: 0}
var VRIGHT = Vec2{X: 1, Y: 0}
var MOTIONS = [4]Vec2{VUP, VDOWN, VLEFT, VRIGHT}
var outMirrors = [4]byte{'^', 'v', '<', '>'}
