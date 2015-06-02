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


type PlainError struct {
	Flag int
	Msg  string	
}

func (e *PlainError) Error() string {
	msgCombined := "state: " +  e.Msg +  "; Flag: "  + string(e.Flag)
	//fmt.Println(msgCombined);
	return msgCombined
}



type Plainer interface {
	Plain() ( float64, *PlainError )
}

func (r Rect) Plain() ( float64, *PlainError) {
	return float64(r.H*r.B), &PlainError{1,"fine1"}
}

func (c Circle) Plain() ( float64, *PlainError) {
	tmp := float64(c.R*c.R) * math.Pi 
	tmp = math.Floor(100*tmp)/100
	return tmp , &PlainError{1,"fine2"}
}

func main() {
	v := Rect{3,5}
	c := Circle{1}
	
	var p Plainer
	p = v
	fmt.Println( p.Plain())

	p = c
	xx, err := p.Plain()
	fmt.Println( xx, err)
	
}



