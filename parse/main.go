package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"

	"github.com/pbberlin/tools/fetch"
	"golang.org/x/net/html"
)

func main() {

	s := `	<p id='1'>Links:
				<span id='11'>p1</span>
				<span id='12'>p2</span>
			</p>
			<ul id='2'>
				<li id='3'><a id='31'  href="foo">Linktext1 <span id='311'>inside</span></a>
				<li id='4'><a id='41'  href="/bar/baz">BarBaz</a>
			</ul>`

	var doc1, doc2 *html.Node
	_, _ = doc1, doc2
	var err error

	doc1, err = html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}

	_, resBytes, err := fetch.UrlGetter("http://localhost:4000/static/handelsblatt.com/article01.html", nil, true)
	doc2, err = html.Parse(bytes.NewReader(resBytes))
	if err != nil {
		log.Fatal(err)
	}

	TraverseHori(Tx{doc1, 0})

	TraverseVert(doc1, 0)

	ioutil.WriteFile("outp.txt", stackOutp, 0)
	// fmt.Printf("\n\n%v", string(resBytes))

}
