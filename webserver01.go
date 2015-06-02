package main

import (
	"fmt"
	"log"
	"net/http"
)

//type AnyType struct{}
type serveHttpAppStruct struct {
	Version, Greeting string
}
type secondHttpStruct string

type thirdHttpStruct struct {
	Greeting, Punct, Who string
}

func (h serveHttpAppStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "1: I say: ", h.Greeting, "|", "Version:", h.Version)
}

func (h secondHttpStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "2: I say: ", h)
}

func (h thirdHttpStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "3: I say: ", h.Greeting, h.Punct, h.Who)
}

func main() {
	http.Handle("/string", secondHttpStruct("I am a String."))

	h3 := thirdHttpStruct{"Hello", " , ", "World"}
	http.Handle("/struct", h3)

	patterns := http.DefaultServeMux.GetPatterns()
	for k, v := range patterns {
		log.Printf("%v %v \n", k, v)
	}

	http.ListenAndServe("localhost:4000", nil)

	/*
		h := serveHttpAppStruct{"1.0", "Hello World of Webservers"}
		http.ListenAndServe("localhost:4000", h)
	*/
}
