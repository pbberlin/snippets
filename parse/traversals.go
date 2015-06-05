package main

import (
	"fmt"

	"github.com/pbberlin/tools/util"
	"golang.org/x/net/html"
)

type Tx struct {
	Nd  *html.Node
	Lvl int
}

var ndStack util.Stack // filled by TraverseVert()
var skipInStack = map[string]bool{"em": true}
var stackOutp []byte

// TraverseHori traverses the tree horizontally.
// It uses a queue. A FiFo structure.
// Inspired by www.geeksforgeeks.org/level-order-tree-traversal/
func TraverseHori(lp interface{}) {

	var queue = util.NewQueue(10)

	lvlPrev := 0
	for lp != nil {

		lpn := lp.(Tx).Nd
		lvl := lp.(Tx).Lvl

		// print current
		if lvl != lvlPrev { // new level => newline
			fmt.Printf("\n%2v:\t", lvl)
			lvlPrev = lvl
		}
		fmt.Printf("%8s  ", lpn.Data)

		// enqueue all children
		for c := lpn.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode {
				queue.EnQueue(Tx{c, lvl + 1})
			}
		}
		lp = queue.DeQueue()
	}
}

func TraverseVert(n *html.Node, lvl int) {

	// Before children
	switch n.Type {
	case html.ElementNode:
		switch n.Data {
		case "a":
			printAttr(n.Attr, []string{"id", "bd"})
		case "iframe", "script", "noscript":
			return
		}

		if !skipInStack[n.Data] {
			ndStack.Push(n.Data)
		}

		// lvl == ndStack.Len()
		s := fmt.Sprintf("%2v: %s\n", ndStack.Len(), ndStack.StringExt(true))
		stackOutp = append(stackOutp, s...) // exceptional comfort case

	case html.TextNode:
		//
	}

	// Children
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		TraverseVert(c, lvl+1)
	}

	// After children
	switch n.Type {
	case html.ElementNode:

		if !skipInStack[n.Data] {
			ndStack.Pop()
		}
	}

}
