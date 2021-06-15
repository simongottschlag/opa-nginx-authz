package main

import (
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
	endpoint       string
}

func (client *handlerClient) opa(w http.ResponseWriter, req *http.Request) {
	opaRequest, err := GetOpaRequest(req, client.endpoint)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create opa request to %q, recevied error: %s\n", client.endpoint, err)
		http.Error(w, "unknown error", http.StatusInternalServerError)
		return
	}

	opaResponse, err := client.httpClient.Do(opaRequest)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to send opa request to %q, recevied error: %s\n", client.endpoint, err)
		http.Error(w, "unable to communicate with upstream", http.StatusInternalServerError)
		return
	}

	defer opaResponse.Body.Close()

	body, err := ioutil.ReadAll(opaResponse.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse response body, recevied error: %s\n", err)
		http.Error(w, "unable to parse response body from upstream", http.StatusInternalServerError)
		return
	}

	var jsonData interface{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to unmarshal json from opa response body, recevied error: %s\n", err)
		http.Error(w, "unable to parse unmarshal json from upstream", http.StatusInternalServerError)
		return
	}

	fmt.Printf("json data: %v\n", jsonData)

	jmespathResult, err := client.jmespathClient.Search(jsonData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse response body with jmespath, recevied error: %s\n", err)
		http.Error(w, "unable to parse response body with jmespath from upstream", http.StatusInternalServerError)
		return
	}

	result, ok := jmespathResult.(bool)
	if !ok {
		fmt.Fprintf(os.Stderr, "Unable to typecast jmespath to bool\n")
		http.Error(w, "unable to typecast jmespath to bool from upstream", http.StatusInternalServerError)
		return
	}

	if !result {
		fmt.Fprintf(os.Stderr, "Result is false\n")
		http.Error(w, "received false result from upstream", http.StatusForbidden)
		return
	}

	fmt.Printf("%s Recevied successful request\n", time.Now())
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Application returned error: %s\n", err)
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
		endpoint:       "http://opa-test:8181/v1/data/nginx/authz",
	}

	http.HandleFunc("/", handlerClient.opa)
	http.ListenAndServe(":8082", nil)

	return nil
}
