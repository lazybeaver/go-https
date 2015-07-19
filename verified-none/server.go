package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s | %s", r.URL, r.UserAgent())
	fmt.Fprintf(w, "%s\n", "Hello World")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("https://localhost:8080/\n")

	if err := http.ListenAndServeTLS(":8080", "../pki/server/certs/localhost.pem", "../pki/server/private/localhost.key", nil); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}
}
