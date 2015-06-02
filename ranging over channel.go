package main

import "fmt"
import "time"
import "math/rand"

/*
	This demonstrates that the ranging
	over a buffered channel
	runs forever.

	It stops at close(channel)
	and not otherwise

*/
func main() {

	rand.Seed(time.Now().UnixNano())

	bufferSize := rand.Intn(3)
	queue := make(chan string, bufferSize)
	fmt.Printf("we buffer with %v \n", bufferSize)

	go func() {
		//time.Sleep(10 * time.Millisecond)
		for elem := range queue {
			fmt.Println(elem)
		}
	}()

	time.Sleep(100 * time.Millisecond)

	queue <- "one"
	queue <- "two"
	queue <- "three"
	queue <- "four"

	time.Sleep(200 * time.Millisecond)

	close(queue)
}
