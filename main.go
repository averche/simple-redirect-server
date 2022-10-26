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
	log.Println("base address:", redirectBaseAddress)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		redirectTo, err := redirectBaseAddress.Parse(r.URL.Path)
		if err != nil {
			log.Printf("%s -> error: %s\n", r.URL.String(), err.Error())
			return
		}

		log.Printf("%s -> %s\n", r.URL.String(), redirectTo.String())

		http.Redirect(w, r, redirectTo.String(), http.StatusMovedPermanently)
	})

	http.ListenAndServe(":9090", nil)
}
