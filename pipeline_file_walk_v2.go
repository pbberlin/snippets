package main



/*
	This is a pipeline with three stages.

	Stage1: ONE go-routine gathers filepaths 
		and sends them to ONE channel (s1c).
		It is NOT many go-routines sending one path to a channel 
		as suggested by John Graham Cumming
			https://www.youtube.com/watch?v=SmoM1InWXr0
			https://github.com/gophercon/2014-talks/blob/master/John_Graham-Cumming_A_Channel_Compendium.pdf
	
	Stage2: SEVERAL go-routines read s1c and process its data, 
		stuffing t into ONE channel (s2c).
		This is called a "bounded" pipeline - because "several" determines
		the number of parallel processing.
		Upstream processing is also limited by 
		the processing routines of Stage2
		
	Stage3:
		ONE go-routine drains s2c and prints its result.
		Again, stage3 limits all previous (upstream) stages,
		because s1c and s2c are unbuffered.
	
	MIGHTITUDEs
		Each stage has an order of parallelity
		Stage1 => 1
		Stage2 => 3
		Stage3 => 1
		
		These "mightitude of parallelity" of course needs to be
		adapted to the cost of processing.
		If stage 1 were more expensive to process, 
		then we would increase the number of go-routines in that stage.
		We adapt so that every downstream stage has slightly more
		capacity then the previous stage - processing is "sucked" 
		downstream - not held up.
	
		
	SIGNALLING
		All go-routines select over the "done" chanel.
		They terminate upon receiving from the done channel.
		Thereby ANY go-routine or any external consumer may
		cause all stages to end.
		
		To send a done token to all stages and all helper go-routines,
		we would have to keep track of their number.
		But we rather CLOSE the done channel - thus
		all receive requests get through at once.
		
		Default it: the first stage closes "done", after finishing
		reading all input-data.
	
		To free resources, the "done" branches also close s1c and s2c.
		This creates the danger of receving "Zero" channel values 
		upstream or downstream.
		Therefore we check for "closedness" with
			val,ok := <- chan  ; if !ok { //chan was closed ; return }
	
	I adapted
		http://blog.golang.org/pipelines
		http://blog.golang.org/pipelines/bounded.go

	Implementation preferences:
	For clarity, I prefer global variables for the major channels,
	instead of creating them inside of functions and passing them around.	
	
	I also cling to making the filepath.Walk helper function explicit,
	instead of using anonymous closure function.
	This robs me of passing the channels as closure arguments:
		closure_arg1 := ...
		closure_arg2 := ...
		filepath.Walk( root, func(...)error{...} )
	
	I prefer  
		filepath.Walk( root, myHelper )


	After calling close, 
	=> all previously sent values will be delivered
	=> then all receive operations will return the 
			ZERO VALUE for the channel's type 
			without blocking. 
			

	We can check for "closedness" with
		val,ok := <- chan
		if !ok { fmt.Println("channel was closed") }
		
	A nil channel is never ready for communication 
	(opposite of closed channel)
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

var s1_eglo  chan error  = make(chan error, 1)   // communicating the - one - return err of walk(...)
var s1_edet  chan error  = make(chan error)      // all small fw errors

var s2c   chan result = make(chan result)

var done  chan struct{}   = make(chan struct{})  // THE one any only termination channel


var funcFW  func(string, os.FileInfo, error) error = func(path string, info os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if !info.Mode().IsRegular() {
		s1_edet <- errors.New("irregular " + path)
		return nil
	}

	select {
		case s1c <- path :
			fmt.Println("\t\t\t loading ",path)
			//if len(path) < 33 {
			//	s1_edet <- errors.New( path + " is under 20")
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
				case   err := <- s1_edet:
					fmt.Println("\t fw detail error: ", err)
				case          <-done:
					fmt.Println("\t s1_edet loop done")
					return
			}
		}
	}()


	go func() {
		fmt.Println("starting fw")
		s1_eglo <- filepath.Walk(root,funcFW)
		protectedClose( s1c ) 		// Close the path channel after Walk returns
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
				case err := <-s1_eglo :
					if err != nil {
						fmt.Println("\t\t\t  s1_eglo loop - err was ", err)
					}
				case      <-done:
					fmt.Println("\t\t\t  s1_eglo loop - done")
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

/*
Wrapper for closing a channel of string
=> Dont panic if channel is nil or already closed
=>	the deferred func catches and neutralizes the panic

Another application of panic-recover is in the json package.
	JSON-data is encoded recursively.
	Malformed JSON-data triggers a panic,
		undwinding the stack to the top-level.
	The top-level function calls recover() and returns a standard error.

This is idiomatic procedure in packages:
Use panic() internally, but return standard errors 
to the external world.
*/

func protectedClose( c chan string ) {

	defer func(){
		//fmt.Println("--protectedClose(): done")  // Println works even in panic
		if panicCond := recover(); panicCond != nil {
			fmt.Printf("PANIC: %v\n", panicCond)
		}
	}()
	
	close(c)
}