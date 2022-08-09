package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <redirect-base-address>\n", os.Args[0])
	}

	redirectBaseAddress, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err.Error())
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		redirectTo := *r.URL
		redirectTo.Host = redirectBaseAddress.Host

		log.Printf("%s -> %s\n", r.URL.String(), redirectTo.String())

		http.Redirect(w, r, redirectTo.String(), http.StatusMovedPermanently)
	})

	http.ListenAndServe(":9090", nil)
}
