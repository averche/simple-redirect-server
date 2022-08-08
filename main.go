package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <redirect-url>\n", os.Args[0])
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, os.Args[1], http.StatusMovedPermanently)
	})

	http.ListenAndServe(":9090", nil)
}
