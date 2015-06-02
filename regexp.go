// Go offers built-in support for [regular expressions](http://en.wikipedia.org/wiki/Regular_expression).
// Here are some examples of  common regexp-related tasks
// in Go.

package main

import "bytes"
import "fmt"
import "regexp"

func main() {


    r, _ := regexp.Compile("p([a-z_]+?)ch")
    s := "This is a patriarch or a peach or a punch or a pinch"

    fmt.Printf ("match:     %v\n" , r.MatchString       (s))
    fmt.Println("find:      " + r.FindString        (s))
    fmt.Printf ("submatch: %v\n",r.FindStringSubmatch(s))
    fmt.Println("findall: ",r.FindAllString     (s,  4) )
    fmt.Println()

    // `[]byte` arguments instead of `String` 
    fmt.Print("byte: ",r.Match([]byte("peach")))

    // When creating constants with regular expressions
    // you can use the `MustCompile` variation of `Compile`. 
    // A plain `Compile` won't work for constants because of 2 return vals.
    r = regexp.MustCompile("p([a-z]+)ch")  ;    fmt.Print("     --  regex const: ",r,"\n")


    // replace - with string or with func
    fmt.Println(r.ReplaceAllString(s, "<fruit>"))
    in := []byte(s)
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}