package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
)

/*
	// negroni extends the http.HandlerFunc concept by one argument;
	//	a next http.HandlerFunc

	type Handler interface {
	        ServeHTTP(ResponseWriter, *Request)
	}
	type HandlerFunc func(ResponseWriter, *Request)
	func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
		f(w, r)
	}

	type Handler interface {
		ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
	}
	type HandlerFunc func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)

	func (h HandlerFunc) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		h(rw, r, next)
	}

	// old fashioned http.Handler are extened and converted:
	func Wrap(handler http.Handler) Handler {
		return HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
			handler.ServeHTTP(rw, r)
			next(rw, r)
		})
	}

*/

var pf = fmt.Printf
var wpf = fmt.Fprintf
var lpf = log.Printf

func Mw1(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.URL.String() == "/favicon.ico" {
		return
	}
	lpf("bef1 %v ", r.URL)
	next(w, r)
}
func Mw2(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.URL.String() == "/favicon.ico" {
		return
	}
	lpf("bef2 %v ", r.URL)
	next(w, r)
}
func Mw3(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.URL.String() == "/favicon.ico" {
		return
	}
	lpf("bef3 %v ", r.URL)
	next(w, r)
	lpf("aft3 %v ", r.URL)
	lpf("\n")
}
func MwA(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.URL.String() == "/favicon.ico" {
		return
	}
	lpf("AdminMW %v ", r.URL)
	next(w, r)
}

func main() {

	Mw1h := negroni.HandlerFunc(Mw1)
	Mw2h := negroni.HandlerFunc(Mw2)
	Mw3h := negroni.HandlerFunc(Mw3)
	MwAh := negroni.HandlerFunc(MwA)

	// Negroni stack of middlewares
	n := negroni.Classic()
	// or
	n = negroni.New(Mw1h, Mw2h) // New() or
	n.Use(Mw3h)                 // Use() => same effect
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery())

	n.Use(negroni.NewStatic(http.Dir("c:\\temp")))

	//
	// Standard routes
	// =====================
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		wpf(w, "\n\nHome page!\n\n")
	})
	mux.HandleFunc("/abc", func(w http.ResponseWriter, req *http.Request) {
		wpf(w, "\n\nabc page!\n\n")
	})

	//
	// Classic "tie in" of another mux with routes:
	muxEditor := http.NewServeMux()
	muxEditor.HandleFunc("/editor/1", func(w http.ResponseWriter, req *http.Request) {
		wpf(w, "\n\nEditor page 1!\n\n")
	})
	muxEditor.HandleFunc("/editor/2", func(w http.ResponseWriter, req *http.Request) {
		wpf(w, "\n\nEditor page 2!\n\n")
		v := 0
		q := 10 / v // causing panic, panic should be caught by middleware
		_ = q
	})
	// Combine mux + muxEditor.
	// Most underdocumented feature I've seen in go.
	// Prefixes must match.
	mux.Handle("/editor/", muxEditor)

	//
	// Now combine main router with muxAdmin router.
	//    And add special middleware.
	muxAdmin := http.NewServeMux()
	muxAdmin.HandleFunc("/admin/", func(w http.ResponseWriter, req *http.Request) {
		wpf(w, "\n\nAdmin home!\n\n")
	})
	muxAdmin.HandleFunc("/admin/1", func(w http.ResponseWriter, req *http.Request) {
		wpf(w, "\n\nAdmin page 1!\n\n")
	})
	muxAdmin.HandleFunc("/admin/2", func(w http.ResponseWriter, req *http.Request) {
		wpf(w, "\n\nAdmin page 2!\n\n")
	})

	//                    \|/ here is the leap in abstraction. Negroni wraps entire mux instances with all included handlers
	mux.Handle("/admin/", negroni.New(MwAh, negroni.Wrap(muxAdmin)))

	// Combine negroni main stack with main router.
	// Another magic heap
	n.UseHandler(mux)

	n.Run(":8091")
}
