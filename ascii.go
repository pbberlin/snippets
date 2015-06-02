package main

import "fmt"

func main() {
	
	
	xx := 'a'		// ascii code
	xx2 := 'A'
	fmt.Printf("%v %v %T  %T\n", xx, xx2, xx, xx2)

	xx++
	xx2++

	fmt.Printf("%c %c \n", xx, xx2)  // ascii to char

	vascii := []byte("x")
	ascii := vascii[0]
	fmt.Printf("\t %d\n", ascii)

}
