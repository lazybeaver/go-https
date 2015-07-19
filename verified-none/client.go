package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Get("https://localhost:8080/")
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
