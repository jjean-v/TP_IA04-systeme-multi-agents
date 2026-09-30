package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func youAre(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You are, Jean")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/jean", youAre)
	s := &http.Server{
		Addr:           ":12000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20}
	log.Println("Listening on localhost:12000")
	log.Fatal(s.ListenAndServe())
}
