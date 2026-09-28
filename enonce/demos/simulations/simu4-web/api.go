package simulation

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func StartAPI(sim *Simulation) {
	mux := http.NewServeMux()

	pif := func(w http.ResponseWriter, r *http.Request) {
		log.Println("PIF!")
		msg, _ := json.Marshal(sim.env.PI())
		fmt.Fprintf(w, "%s", msg)
	}

	mux.HandleFunc("GET /pi", pif)

	s := &http.Server{
		Addr:           ":12000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20}

	log.Println("Listening on localhost:12000")
	log.Fatal(s.ListenAndServe())
}
