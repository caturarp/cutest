package main

// create server with http package
import (
	"fmt"
	"net/http"
)

func main() {
	// create a simple handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	// create a html page
	http.HandleFunc("/html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// start server
	http.ListenAndServe(":8888", nil)

}
