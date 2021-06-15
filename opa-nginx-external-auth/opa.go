package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmespath/go-jmespath"
)

type Input struct {
	Method     string              `json:"method"`
	Body       string              `json:"body"`
	ParsedBody interface{}         `json:"parsed_body"`
	Path       string              `json:"path"`
	Version    string              `json:"version"`
	Headers    map[string][]string `json:"headers"`
}
type OpaInput struct {
	Input `json:"input"`
}

func GetOpaInputJson(req *http.Request) ([]byte, error) {
	body, err := GetStringFromBody(req.Body)
	if err != nil && err.Error() != "Body is nil" {
		return nil, fmt.Errorf("Request body error: %w", err)
	}

	var parsedBody interface{}
	parsedBody = ""
	if body != "" {
		parsedBody = GetStructFromBody([]byte(body))
	}

	path := ""
	if req.URL != nil {
		path = req.URL.Path
	}

	receivedRequest := Input{
		Method:     req.Method,
		Body:       body,
		ParsedBody: parsedBody,
		Path:       path,
		Version:    fmt.Sprintf("%d.%d", req.ProtoMajor, req.ProtoMinor),
		Headers:    req.Header,
	}

	return json.Marshal(OpaInput{receivedRequest})
}

func GetOpaRequest(req *http.Request, endpoint string) (*http.Request, error) {
	opaInput, err := GetOpaInputJson(req)
	if err != nil {
		return nil, err
	}

	opaRequest, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(opaInput))
	if err != nil {
		return nil, err
	}

	opaRequest.Header.Add("Content-Type", "application/json")

	return opaRequest, nil
}

func GetOpaHttpResponse(httpClient *http.Client, req *http.Request, endpoint string) (*http.Response, error) {
	opaRequest, err := GetOpaRequest(req, endpoint)
	if err != nil {
		return nil, err
	}

	opaResponse, err := httpClient.Do(opaRequest)
	if err != nil {
		return nil, err
	}

	return opaResponse, nil
}

func GetOpaResponseStruct(httpClient *http.Client, req *http.Request, endpoint string) (interface{}, error) {
	opaResponse, err := GetOpaHttpResponse(httpClient, req, endpoint)
	if err != nil {
		return nil, err
	}

	body, err := GetStringFromBody(opaResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("Response body error: %w", err)
	}

	var jsonData interface{}
	err = json.Unmarshal([]byte(body), &jsonData)
	if err != nil {
		return nil, fmt.Errorf("Unable to unmarshal OPA response: %w", err)
	}

	return jsonData, nil
}

func GetResultFromOpaResponseStruct(client *jmespath.JMESPath, response interface{}) (bool, error) {
	jmespathResult, err := client.Search(response)
	if err != nil {
		return false, err
	}

	result, ok := jmespathResult.(bool)
	if !ok {
		return false, err
	}

	return result, nil
}
