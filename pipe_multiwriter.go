package main

import (
	"fmt"
	"io"
	"os"
	"time"
)



func main() {

	r, w := io.Pipe() 

	// reading	
	go func(){
		
		for  i:=0; i< 10;i++{
			var b []byte
			b = make([]byte,10)	
			n,err := r.Read(b)
			if err != nil {
				fmt.Printf("lp %v: - err -%v- \n",i, err)			
				return
			}
			fmt.Printf("\t\t  lp %v: bytes read %v \n",i, n)			
			fmt.Printf("\t\t    %s\n", b)			
		}
	}()	


	// writing
	go func(){
		
		mw := io.MultiWriter(os.Stdout,w)
		_ = mw
		fmt.Fprintf(mw,"hello\n")
		fmt.Printf("Printed to pipe 1\n")		

		fmt.Fprintf(mw,"we write to the pipe writer\n")
		fmt.Printf("Printed to pipe 2\n")		
	}()
	
	time.Sleep(500* time.Millisecond)	

	
	
}

