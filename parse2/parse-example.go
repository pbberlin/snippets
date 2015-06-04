package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/pbberlin/tools/util"
	"golang.org/x/net/html"
)

var f func(*html.Node)
var stack util.Stack

func init() {
	f = func(n *html.Node) {
		switch n.Type {
		case html.ElementNode:
			// if n.Data == "a" || true {
			// 	printAttr(n.Attr, []string{"id", "bd"})
			// }
			stack.Push(n.Data)
			fmt.Printf("%2v %v\n", stack.Len(), stack)

		case html.TextNode:
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}

		switch n.Type {
		case html.ElementNode:
			stack.Pop()
		}

	}
}

func main() {

	s := `	<p id='1'>Links:</p>
			<ul id='2'>
				<li id='3'><a id='31'  href="foo">Linktext1 <span id='311'>inside</span></a>
				<li id='4'><a id='41'  href="/bar/baz">BarBaz</a>
			</ul>`

	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}

	f(doc)

}
