package main

import( 
	"fmt"
	"reflect"
)


func main() {

	type MyInt int

	var a  int=4
	var b  MyInt=4
	var c  interface{}
	
	fmt.Printf("Types: %3v %11v - %18v\n", reflect.TypeOf(a), reflect.TypeOf(c), reflect.ValueOf(c))

	c = a
	cv := reflect.ValueOf(c)
	ci := cv.Interface()
	fmt.Println(ci)
	fmt.Printf("Types: %3v %11v - %18v\n", reflect.TypeOf(a), reflect.TypeOf(c), reflect.ValueOf(c))


	c = b
	cv = reflect.ValueOf(c)
	ci = cv.Interface()
	fmt.Println(ci)
	fmt.Printf("Types: %3v %11v - %18v\n", reflect.TypeOf(a), reflect.TypeOf(c), reflect.ValueOf(c))


}