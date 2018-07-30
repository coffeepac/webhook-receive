package main

import (
	"fmt"
	http "net/http"
	"time"
)

func handleWebHooks(w http.ResponseWriter, r *http.Request) {
	// let nexmo know we're rad
	fmt.Fprintf(w, "receipt recieved")

	// dump the info
	fmt.Println(r)
	bytes := make([]byte, 200)
	r.Body.Read(bytes)
	fmt.Println(string(bytes))
}

func main() {
	http.HandleFunc("/", handleWebHooks)

	s := &http.Server{
		Addr:         ":80",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	s.ListenAndServe()
}
