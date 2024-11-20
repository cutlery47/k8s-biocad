package main

import (
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func main() {
	http.HandleFunc("/", helloHandler)

	log.Println("Server is running")
	log.Fatal(http.ListenAndServe(":32777", nil))
}
