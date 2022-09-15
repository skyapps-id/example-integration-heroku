package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	port := "3000"
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	http.HandleFunc("/", helloHandler)
	log.Println("Listing for " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
