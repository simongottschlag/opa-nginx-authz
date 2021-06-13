package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/jmespath/go-jmespath"
)

type handlerClient struct {
	httpClient     *http.Client
	jmespathClient *jmespath.JMESPath
}

func (client *handlerClient) opa(w http.ResponseWriter, req *http.Request) {
	data, err := json.Marshal(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to marshal json from req: %s", err)
		http.Error(w, "unknown error", http.StatusInternalServerError)
		return
	}

	endpoint := "http://opa-test:8181/v1/data/nginx/authz"

	opaRequest, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(data))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create request: %s", err)
		http.Error(w, "unknown error", http.StatusInternalServerError)
		return
	}

	opaResponse, err := client.httpClient.Do(opaRequest)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to send request to %q, recevied error: %s", endpoint, err)
		http.Error(w, "unable to communicate with upstream", http.StatusInternalServerError)
		return
	}

	defer opaResponse.Body.Close()

	body, err := ioutil.ReadAll(opaResponse.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse response body, recevied error: %s", err)
		http.Error(w, "unable to parse response body from upstream", http.StatusInternalServerError)
		return
	}

	jmespathResult, err := client.jmespathClient.Search(string(body))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse response body with jmespath, recevied error: %s", err)
		http.Error(w, "unable to parse response body with jmespath from upstream", http.StatusInternalServerError)
		return
	}

	result, ok := jmespathResult.(bool)
	if !ok {
		fmt.Fprintf(os.Stderr, "Unable to typecast jmespath to bool")
		http.Error(w, "unable to typecast jmespath to bool from upstream", http.StatusInternalServerError)
		return
	}

	if !result {
		http.Error(w, "received false result from upstream", http.StatusForbidden)
		return
	}
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Application returned error: %s", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func run() error {
	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
		},
		Timeout: 10 * time.Second,
	}

	jmsepathClient, err := jmespath.Compile("result")
	if err != nil {
		return err
	}

	handlerClient := &handlerClient{
		httpClient:     httpClient,
		jmespathClient: jmsepathClient,
	}

	http.HandleFunc("/", handlerClient.opa)
	http.ListenAndServe(":8082", nil)

	return nil
}
