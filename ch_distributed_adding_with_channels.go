package main

import "fmt"
import "time"
import "runtime"


func sum(a []int, sleep time.Duration , signal int) {
    sum := 0
    for _, v := range a {
        sum += v
    }
    fmt.Println( "stt",1000 * sleep * time.Millisecond )
    time.Sleep(  1000 * sleep * time.Millisecond )
    c <- sum // send sum to c
    fmt.Print("snt ",sleep,"\n")
    if signal == 1 {
    	cs <- 1
		fmt.Print("finalize signal sent","\n") 	
    }
}

/*
	ever channel has a default buffer of ONE value
	push onto an empty channel goes through
	push again blocks - until someone has pulled 
	pull from an empty channels blocks, until someone has pushed
		something onto i

*/

var   cs  = make(chan int)
var    c  = make(chan int,2)

func main() {
    a := []int{7, 2, 8, -9, 4, 0,3,4,5}
    chunks := 3
    
    procs := runtime.GOMAXPROCS(4)
    fmt.Println("processes: ", procs)


	for i:=0 ; i<chunks; i++ {
	    go sum(a[i*len(a)/chunks:(i+1)*len(a)/chunks], time.Duration(3*i+2), (i+1)/chunks)	
	}
	sig := <- cs
   x, y, z := <-c, <-c, <-c // receive from c

    fmt.Println(sig, x, y,z,"the sum is", x+y+z)
}