package main

// http://blog.golang.org/pipelines

/*
	After calling close, 
	and after any previously sent values have been received, 
	receive operations will return the zero value for the channel's type 
	without blocking. 

*/

import "fmt"
import "os"
import "errors"
import "sync"
import "path/filepath"
import "time"

type result struct {
	path string
	plen int
}

var s1c   chan string = make(chan string)	    // loaded with paths to process by filewalk

var eglob chan error  = make(chan error, 1)   // communicating the - one - return err of walk(...)
var efw   chan error  = make(chan error)      // all small fw errors

var s2c   chan result = make(chan result)


var done  chan struct{}   = make(chan struct{})


var funcFW  func(string, os.FileInfo, error) error = func(path string, info os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if !info.Mode().IsRegular() {
		efw <- errors.New("irregular " + path)
		return nil
	}

	select {
		case s1c <- path :
			fmt.Println("\t\t\t loading ",path)
			//if len(path) < 33 {
			//	efw <- errors.New( path + " is under 20")
			//}
		case	    <- done :
			return errors.New("file walk canceled")
	}
	return nil
}

func walkFiles(root string) {

	// catch the current messages from file walk
	go func() {
		for {
			select {
				case   err := <- efw:
					fmt.Println("\t fw detail error: ", err)
				case          <-done:
					fmt.Println("\t efw loop done")
					return
			}
		}
	}()


	go func() {
		fmt.Println("starting fw")
		eglob <- filepath.Walk(root,funcFW)
		close(s1c)  		// Close the path channel after Walk returns.
		fmt.Println("ending fw")
	}()
}


func digester(idx int, wg *sync.WaitGroup){
	go func() {
		//cntr := 0; 
		fmt.Println("\t\tdigester lp",idx," start")
		for {
			//cntr++ ; if cntr > 44 { return}
			select {
				case  l,ok := <-s1c:
					if ok {
						s2c <- result{l, len(l)}
						fmt.Println("\t\tdigester lp",idx," - path read: -"+l+"-")						
					} else {
						// upstream s1c may got closed
					}
				case      <-done:
					wg.Done()
					fmt.Println("\t\tdigester lp",idx," done")
					return
			}
		}
	}()
}


func main(){
	
	
	
	walkFiles(`c:\TEMP\`)


	
	const numDigesters = 3

	var wg sync.WaitGroup
	wg.Add(numDigesters)	

	for i := 0; i < numDigesters; i++ {
		digester(i,&wg)
	}


	// after all digesters have finished
	// => close stage two channel
	go func() {
		wg.Wait()
		close(s2c)
		fmt.Println("S2C CLOSED")
	}()



	// summing up
	go func() {
		sum := 0
		fmt.Println("\t\t   start summer up ...")
		LabelX:
		for  {
			select {
				case r := <- s2c:
					sum += r.plen
					fmt.Println("\t\t\t   summed -" + r.path + "-  ", sum)
				case      <-done:
					fmt.Println("\t\t\t   summing done")
					break LabelX
			}
		}
		fmt.Println("\t\t   Sum of all lengths is ",sum)
	}()


	time.Sleep( 451 * time.Millisecond)	

	
	
	// Check whether the Walk failed
	go func() {
		LabelY:
		for  {
			select {
				case err := <-eglob :
					if err != nil {
						fmt.Println("\t\t\t  eglob loop - err was ", err)
					}
				case      <-done:
					fmt.Println("\t\t\t  eglob loop - done")
					break LabelY
			}
		}
	}()


	//done <- struct{}{}
	//done <- struct{}{}	
	close(done)

	time.Sleep( 222 * time.Millisecond)	

	fmt.Println("end")
}