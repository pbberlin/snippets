package main

import "fmt"
import "time"
import "math/rand"

// from
// http://talks.golang.org/2012/concurrency.slide#43

type Result string
type Search func(query string) Result

func getSearchFunc(kind string) Search {
	return func(query string) Result {
		nms := rand.Intn(100)
		time.Sleep(time.Duration(nms) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q %v ms", kind, query, nms))
	}
}

func First(modePike bool, query string, replicas ...Search) Result {
	c := make(chan Result)

	if modePike {
		searchReplica := func(i int) { c <- replicas[i](query) }
		for i := range replicas {
			go searchReplica(i)
		}

	} else {
		// the closure leads to i being evaluated 
		// after the for loop has finished 
		// => its always "max"
		for i := range replicas {
			go func() { c <- replicas[i](query) }()
		}

	}
	return <-c
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Mode Pike results are stochastic:")
	for i := 0; i < 5; i++ {
		result := First(true , "golang", getSearchFunc("w1"), getSearchFunc("w2"), getSearchFunc("w3"), getSearchFunc("w4"))
		fmt.Println(" ",result)
	}

	fmt.Println("Mode Mine - its always the the last replica:")
	for i := 0; i < 5; i++ {
		result := First(false, "golang", getSearchFunc("w1"), getSearchFunc("w2"), getSearchFunc("w3"), getSearchFunc("w4"))
		fmt.Println(" ",result)
	}

}
