package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	hostname, _ := os.Hostname()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World from Go! 🙀🙀🙀 %s\n", hostname)
		fmt.Println("GET " + r.URL.Path)
	})

	fmt.Println("Server started.")
	http.ListenAndServe(":8000", nil)
}
