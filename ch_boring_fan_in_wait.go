package main

import "fmt"
import "time"
import "math/rand"



// from http://talks.golang.org/2012/concurrency.slide#27

func boring(msg string) <-chan Message { // Returns receive-only channel of strings.
    c := make(chan Message)
    waitForIt := make(chan bool) // Shared between all messages.
    go func() { // We launch the goroutine from inside the function.
        for i := 0; ; i++ {
            c <- Message{ fmt.Sprintf("%s: %d", msg, i), waitForIt }
            time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
            <-waitForIt 
        }
    }()
    return c // Return the channel to the caller.
}

func fanIn(input1, input2 <-chan Message) <-chan Message{
    c := make(chan Message)
    go func() { for { c <- <-input1 } }()
    go func() { for { c <- <-input2 } }()
    return c
}


func main() {
    c := fanIn(boring("Joe"), boring("Ann"))
    for i := 0; i < 5; i++ {
        msg1 := <-c; fmt.Println(msg1.str)
        msg2 := <-c; fmt.Println(msg2.str)
        msg1.wait <- true
        msg2.wait <- true
    }    
    fmt.Println("You're both boring; I'm leaving.")
}

type Message struct {
    str string
    wait chan bool
}
