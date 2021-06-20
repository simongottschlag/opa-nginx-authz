package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jmespath/go-jmespath"
)

type HandlerClient struct {
	httpClient     *http.Client
	jmespathClient *jmespath.JMESPath
	opaClient      *OpaClient
	endpoint       string
}

func NewHandlerClient(httpClient *http.Client, jmsepathClient *jmespath.JMESPath, opaClient *OpaClient, endpoint string) *HandlerClient {
	return &HandlerClient{
		httpClient:     httpClient,
		jmespathClient: jmsepathClient,
		endpoint:       endpoint,
		opaClient:      opaClient,
	}
}

func (client *HandlerClient) OpaProxyHandler(w http.ResponseWriter, req *http.Request) {
	opaResponse, err := GetOpaResponseStruct(client.httpClient, req, client.endpoint)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get OPA response: %s\n", err)
		http.Error(w, "unable to communicate with upstream", http.StatusInternalServerError)
		return
	}

	result, err := GetResultFromOpaResponseStruct(client.jmespathClient, opaResponse)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get result: %s\n", err)
		http.Error(w, "unable to get result", http.StatusInternalServerError)
		return
	}

	if !result {
		fmt.Fprintf(os.Stderr, "Result is false\n")
		http.Error(w, "received false result from upstream", http.StatusForbidden)
		return
	}
}

func (client *HandlerClient) OpaRegoHandler(w http.ResponseWriter, req *http.Request) {
	input, err := GetOpaInput(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to generate opa input: %s\n", err)
		http.Error(w, "unable to generate opa input", http.StatusInternalServerError)
		return
	}

	result, err := GetResultWithOpaInput(req.Context(), client.opaClient, input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get opa result: %s\n", err)
		http.Error(w, "unable to get opa result", http.StatusInternalServerError)
		return
	}

	if !result {
		fmt.Fprintf(os.Stderr, "Result is false\n")
		http.Error(w, "result is false", http.StatusForbidden)
		return
	}
}
