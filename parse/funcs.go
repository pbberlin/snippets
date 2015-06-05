package main

import (
	"fmt"

	"golang.org/x/net/html"
)

func printAttr(attributes []html.Attribute, keys []string) {
	for _, a := range attributes {
		for i := 0; i < len(keys); i++ {
			if keys[i] == a.Key {
				fmt.Printf("id is %v\n", a.Val)
			}
		}
	}
}

func printLvl(n *html.Node, col int) {
	if n.Type == html.ElementNode {
		fmt.Printf("%2v: %2v ", col, n.Data)
	}

}
