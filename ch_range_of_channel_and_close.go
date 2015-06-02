package main

import "fmt"
/*
comment in the "close(strings)"
and "range" waits forever

it appears, s:= range strings is equivalent
for{ s,ok:= <- strings;  if !ok {break} }
*/


func generator(strings chan string) {
	strings <- "Five hour's New York jet lag"
	strings <- "to the dire and ever-decreasing circles"
	strings <- "of disrupted rhythm."
	close(strings)
}

func main() {
	strings := make(chan string)
	go generator(strings)
	for s := range strings {
		fmt.Printf("%s \n", s)
	}
	fmt.Printf("END\n")
}