package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s | %s", r.URL, r.UserAgent())
	fmt.Fprintf(w, "%s\n", "Hello World")
}

func main() {
	pool := x509.NewCertPool()
	cert, err := ioutil.ReadFile("../pki/ca/certs/localhost.pem")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}
	pool.AppendCertsFromPEM(cert)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	fmt.Printf("https://localhost:8080/\n")
	if err := server.ListenAndServeTLS("../pki/server/certs/localhost.pem", "../pki/server/private/localhost.key"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}
}
