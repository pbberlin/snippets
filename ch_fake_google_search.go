package main

import "fmt"
import "time"
import "math/rand"

// from http://talks.golang.org/2012/concurrency.slide#43


type Result string
type Search func(query string) Result


func getSearchFunc(kind string) Search {
		return func(query string) Result {
			ts := rand.Intn(100)
			time.Sleep(time.Duration(ts) * time.Millisecond)
			return Result(fmt.Sprintf("%s result for %q %v\n", kind, query, ts))
		}
}

var (
	webSearcher = getSearchFunc("web")
	imgSearcher = getSearchFunc("image")
	vidSearcher = getSearchFunc("video")

	webSearchr1 = getSearchFunc("web_1")
	webSearchr2 = getSearchFunc("web_2")
	webSearchr3 = getSearchFunc("web_3")
	webSearchr4 = getSearchFunc("web_4")

	dummy = getSearchFunc("searchdum")

)

// succinctly
func Google1(query string) (results []Result) {
	results = append(results, webSearcher(query)  )
	results = append(results, imgSearcher(query))
	results = append(results, vidSearcher(query))
	return
}

// parallel
func Google2(query string) (results []Result) {

	c := make(chan Result)
	go func() { c <- webSearcher(query) } ()
	go func() { c <- imgSearcher(query) } ()
	go func() { c <- vidSearcher(query) } ()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}


// parallel, max 60 ms
func Google3(query string) (results []Result) {

	c := make(chan Result)
	go func() { c <- webSearcher(query) } ()
	go func() { c <- imgSearcher(query) } ()
	go func() { c <- vidSearcher(query) } ()

	timeout := time.After(60 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("some requests timed out")
			return
		}
	}
	return
}

// http://talks.golang.org/2012/concurrency.slide#48
// returning the FIRST response from several web/img/vid servers
func First(query string, replicas ...Search) Result {
	c := make(chan Result)

/*


	for i := 0; i < len(replicas)-1 ; i++	{
		fmt.Printf("%v %+#v\n",i, replicas[i])		
		go func(){ c <- replicas[i](query) }()
	}

	go func(){ c <- replicas[0](query) }()
	go func(){ c <- replicas[1](query) }()
	go func(){ c <- replicas[2](query) }()
	go func(){ c <- replicas[3](query) }()


	for i,r := range replicas {
		fmt.Printf("%v %#v\n",i, r)		
		go func(){ c <- r(query) }()
	}

	searchReplica := func(i int) { c <- replicas[i](query) }	
	for i := range replicas {
		fmt.Printf("%v \n",i)		
		go searchReplica(i)
	}	

	searchReplica := func(i int) { c <- replicas[i](query) }	
	for i := range replicas {
		go searchReplica(i)
	}	

	for i := range replicas {
		go func(){ c <- replicas[i](query) }()
	}

*/	

	searchReplica := func(i int) { c <- replicas[i](query) }	
	for i := range replicas {
		go searchReplica(i)
	}	



	return <-c
}

func main1(){
	start := time.Now()
	results := Google3("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)	
}

func main2() {
	start  := time.Now()
	//result :=	 First("golang" , webSearchr1,webSearchr2, webSearchr3, webSearchr4)
	result :=	 First("golang" , getSearchFunc("w1"),getSearchFunc("w2"), getSearchFunc("w3"), getSearchFunc("w4") )
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)
}

func main(){
	rand.Seed(time.Now().UnixNano())

	main2()
}
