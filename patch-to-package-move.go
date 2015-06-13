package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// for pkg := range subpackages(ctxt, srcDir, from) {
	//   :
	//   :
	//   :

	// filepath.Join ensures normalized paths with trailing "/".
	_, ffr := filepath.Split(filepath.Join(from, ""))
	_, fto := filepath.Split(filepath.Join(to, ""))
	destinations[pkg] = strings.Replace(filepath.Join(from, ""), ffr, fto, 1)
	fmt.Printf("1: %v\n2: %v\n3: %v\n", destinations[pkg], ffr, fto)

	// }
}
