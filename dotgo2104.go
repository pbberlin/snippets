package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

type task interface {
	perform(i int) string
	print()
}

type t1 struct{ s string }
type t2 struct{ s string }

func (x t1) print() {
	fmt.Println("t1", x.s)
}
func (x t2) print() {
	fmt.Println("t2", strings.Repeat(x.s+"-", 4))
}

func (x t1) perform(i int) string {
	return fmt.Sprint("t1 processed by ", i, " ", x.s)
}
func (x t2) perform(i int) string {
	return fmt.Sprint("t2 processed by ", i, " ", x.s)
}

var in = make(chan task)
var out = make(chan string)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("  --<")
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			fmt.Println("  --" + s.Text())
			time.Sleep(10 * time.Millisecond)
			if s.Text() == "quit" {
				break
			}
			if rand.Intn(1e3)%2 == 0 {
				in <- t2{s: s.Text()}

			} else {
				in <- t1{s: s.Text()}

			}
		}
		wg.Done()

	}()

	// process
	for i := 0; i < 4; i++ {
		wg.Add(1)
		j := i
		go func(i int) {
			for {
				t := <-in
				res := t.perform(i)
				out <- res
			}
			wg.Done()
		}(j)
	}

	//
	// print
	wg.Add(1)
	go func() {
		for {
			res := <-out
			fmt.Println(res)
		}
		wg.Done()
	}()

	//
	wg.Wait()
	fmt.Println("main end")
}

func init() {
	rand.Seed(time.Now().UnixNano())

}
