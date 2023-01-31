package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, _ *http.Request){
	fmt.Fprintf(w, "Hello world!");
}

func main() {
	fmt.Printf("Hello, world!")

	http.HandleFunc("/", home);

	http.ListenAndServe(":3333", nil)
}