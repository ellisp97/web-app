package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Health Check</h1>")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><h2>Hey there, thanks for entering %s!</h2></html>", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
