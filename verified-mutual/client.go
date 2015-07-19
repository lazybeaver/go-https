package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	pool := x509.NewCertPool()
	caCert, err := ioutil.ReadFile("../pki/ca/certs/localhost.pem")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}
	pool.AppendCertsFromPEM(caCert)

	clientCert, err := tls.LoadX509KeyPair("../pki/client/certs/localhost.pem", "../pki/client/private/localhost.key")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      pool,
				Certificates: []tls.Certificate{clientCert},
			},
		},
	}
	resp, err := client.Get("https://localhost:8080")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}
	defer resp.Body.Close()
	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	} else {
		fmt.Print(string(body))
	}
}
