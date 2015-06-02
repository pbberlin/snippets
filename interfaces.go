package main

import (
	"fmt"
	"math"
)

type Rect struct {
	H,B int
}

type Circle struct {
	R int
}

type Plainer interface {
	Plain() float64
}

func (r Rect) Plain() float64 {
	return float64(r.H*r.B)
}

func (c Circle) Plain() float64 {
	tmp := float64(c.R*c.R) * math.Pi 
	tmp = math.Floor(100*tmp)/100
	return tmp
}

func main() {
	v := Rect{3,5}
	c := Circle{1}
	
	var p Plainer
	p = v
	fmt.Println( p.Plain())

	p = c
	fmt.Println( p.Plain())
	
}