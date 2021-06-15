package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Application returned error: %s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func run() error {
	httpClient := NewHttpClient()

	jmsepathClient, err := NewJmsepathClient("result")
	if err != nil {
		return err
	}

	handlerClient := NewHandlerClient(httpClient, jmsepathClient, "http://opa-test:8181/v1/data/nginx/authz")

	http.HandleFunc("/", handlerClient.OpaHandler)
	http.ListenAndServe(":8082", nil)

	return nil
}
