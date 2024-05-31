package main

// create server with http package
import (
	"fmt"
	"net/http"
)

func main() {
	// create a simple handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, eid 1444H!")
	})

	// create a html page
	http.HandleFunc("/html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	// create log if server up
	fmt.Println("Server is running on port :8000")
	// start server
	http.ListenAndServe(":8000", nil)

}
