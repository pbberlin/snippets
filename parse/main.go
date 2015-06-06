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

	s := `	<p>Links:
				<span>p1</span>
				<span>p2</span>
			</p>
			<ul>
				<li id='332' ><a   href="foo">Linktext1 <span>inside</span></a>
				<li><a   href="/bar/baz">BarBaz</a>
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

	// to file
	var b bytes.Buffer
	err = html.Render(&b, doc1)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("outp.html", b.Bytes(), 0)

}
