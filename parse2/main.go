package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/pbberlin/tools/util"
	"golang.org/x/net/html"
)

var f func(*html.Node, int)
var stack util.Stack

type Tx struct {
	Nd  *html.Node
	Lvl int
}

var queue = util.NewQueue(10)

func init() {
	f = func(n *html.Node, lvl int) {

		// Before children
		switch n.Type {
		case html.ElementNode:
			// if n.Data == "a" || true {
			// 	printAttr(n.Attr, []string{"id", "bd"})
			// }
			stack.Push(n.Data)
			fmt.Printf("%2v: %v\n", stack.Len(), stack) // lvl == stack.Len()

		case html.TextNode:
			//
		}

		// Children
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, lvl+1)
		}

		// After children
		switch n.Type {
		case html.ElementNode:
			stack.Pop()
		}

	}
}

func main() {

	s := `	<p id='1'>Links:
				<span id='11'>p1</span>
				<span id='12'>p2</span>
			</p>
			<ul id='2'>
				<li id='3'><a id='31'  href="foo">Linktext1 <span id='311'>inside</span></a>
				<li id='4'><a id='41'  href="/bar/baz">BarBaz</a>
			</ul>`

	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}

	// f(doc, 0)
	// printLevelwise1()

	PrintByLevel(Tx{doc, 0})

}

// PrintByLevel traverses the tree horizontally.
// It uses a queue. A FiFo structure.
// Inspired by www.geeksforgeeks.org/level-order-tree-traversal/
func PrintByLevel(lp interface{}) {

	lvlPrev := 0
	for lp != nil {

		// print current
		lpn := lp.(Tx).Nd
		lvl := lp.(Tx).Lvl

		if lvl != lvlPrev {
			fmt.Printf("\n%v\t", lvl)
			lvlPrev = lvl
		}
		fmt.Printf("%8s  ", lpn.Data)

		// enqueue all children
		for c := lpn.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode {
				queue.Push(Tx{c, lvl + 1})
			}
		}
		lp = queue.Pop()
	}
}
