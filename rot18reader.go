package main

// http://tour.golang.org/#60

import (
    "io"
    "os"
    "strings"
    "errors"
    "fmt"
    "math"
)

type rot13Reader struct {
    r io.Reader
}


func (rr *rot13Reader) Read(p []byte) (n int, err error){
    
    if len(p) == 0 {
        fmt.Println("no length to store data")
        return 0, nil
    }
    

    var str string = "AZazÄ Hello"
	unicodeCodePoints := []rune(str)
    unicodeCodePoints[  len(unicodeCodePoints) -1 ] = 77
    //fmt.Println( unicodeCodePoints[0], unicodeCodePoints[1], unicodeCodePoints[2], unicodeCodePoints[3], unicodeCodePoints[4] )

    const bsize = 4
    
    sr := (rr.r)
    ra := make( []byte, bsize)

    countTotal := 0
    count := 1
    
    for count > 0 {
	    
        count, err := sr.Read( ra )
        
        if( err != nil ){
	        err = errors.New("error 1")
            //fmt.Println("Error (io.EOF)", err, count , countTotal)
            return count, err
        }
        
        
        for k,v := range ra {
            switch {

                case v >= 65 && v <= 90 : {
                            v += 13
                            if v > 90 {
                                v -= 26
                            }
                            ra[k] = v
                    }
    
    
                case v >= 97 && v <= 122: {
                    v += 13
                    if v > 122 {
                        v -= 26
                    }
                    ra[k] = v
                }

           }
            
        }
        
	    copy( p[countTotal:count], ra[0:count])    
        
        countTotal += count

        effectiveSize := int(math.Min(bsize, float64(count))) 
        effectiveSize++
        effectiveSize--
        return effectiveSize ,nil

        fmt.Println(count, countTotal, p)
    } 
    

	return countTotal, nil

}

func main() {
    
    
    s := strings.NewReader("Lbh penpxrq gur pbqr!")    
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}