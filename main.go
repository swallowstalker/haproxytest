package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	http.HandleFunc("/testhaproxy", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("hello from %s\n", port)
		fmt.Fprintln(w, "hello from "+port)
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("hello from %s\n", port)
		fmt.Fprintln(w, "hello from "+port)
	})

	http.ListenAndServe("0.0.0.0:"+port, nil)
}
