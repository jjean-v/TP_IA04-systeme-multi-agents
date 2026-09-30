package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

type MyServer struct {
	RequestCount int
}

func (s *MyServer) hello(w http.ResponseWriter, r *http.Request) {
	s.RequestCount++
	if r.PathValue("name") == "" {
		fmt.Fprintf(w, "Hello, %q (%d)",
			html.EscapeString(r.URL.Path),
			s.RequestCount)
	} else {
		name := r.PathValue("name")
		fmt.Fprintf(w, "Hello, %q (%d) your name is %s",
			html.EscapeString(r.URL.Path),
			s.RequestCount,
			name)
	}
}

func main() {
	ms := new(MyServer)
	mux := http.NewServeMux()
	mux.HandleFunc("/hello/{name}", ms.hello)
	mux.HandleFunc("/hello/", ms.hello)

	s := &http.Server{
		Addr:           ":12000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20}
	log.Println("Listening on localhost:12000")
	log.Fatal(s.ListenAndServe())
}
