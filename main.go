// The smallest web server with go.
// Usage:
//  go run main.go
// Web browser:
//  http://localhost:8080
package main

import (
	"net/http"
	"fmt"
	"os"
)

func main() {
	// Register root function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK!")
	})

	// Start web server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


