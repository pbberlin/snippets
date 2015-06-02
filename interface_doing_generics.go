package main

import( 
	"fmt"
	_ "time"
	"strconv"
)

// type string or integer
type WrapInt int
type WrapStr string


// interface string or integer
type IWrapStrOrInt interface {
	GetStr() (s string)
	GetInt() (i int)
}


// the type WrapInt and WrapStr 
// now implement IWrapStrOrInt
func (r WrapInt) GetStr() string {
    return 	fmt.Sprintf("%v", r)
}
func (r WrapStr) GetStr() string {
    return 	string(r)
}
func (r WrapInt) GetInt() int{
    return 	int(r)
}
func (r WrapStr) GetInt() int{
    i, _ := strconv.Atoi(string(r))
    return i
}



func main() {

	/*
	    we can now store two different types
	    in one variable:
	*/
	var iorstr IWrapStrOrInt

	iorstr = WrapInt(2)  // assigning an int
	fmt.Printf( "iorstr.GetStr() = %v \ttype %T  \n", iorstr.GetStr() , iorstr.GetStr() )
	fmt.Printf( "iorstr.GetInt() = %v \ttype %T  \n", iorstr.GetInt(),  iorstr.GetInt() )

	
	iorstr = WrapStr("two")   // assigning a string
	fmt.Printf( "iorstr.GetStr() = %v \ttype %T  \n", iorstr.GetStr() , iorstr.GetStr() )
	fmt.Printf( "iorstr.GetInt() = %v \ttype %T  \n", iorstr.GetInt() , iorstr.GetInt() )



	fmt.Println()
	
   // empty interface
	// stores any type
	var eif interface{}
	eif  = 2
	fmt.Printf( "val is %08v  \ttype %T  \n", eif, eif )
	eif  = "two"
	fmt.Printf( "val is %-08v \ttype %T  \n", eif, eif )

}