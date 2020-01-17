package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[GO WEB] [INDEX] Response: Hello World")
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	fmt.Println("Web")

	http.HandleFunc("/", index)

	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", nil)

}
