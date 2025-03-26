package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting Chess Game API on port 8080...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Chess Game API is running")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
