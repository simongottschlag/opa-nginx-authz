package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func RequestBodyToString(req *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	closeErr := req.Body.Close()
	if closeErr != nil {
		return nil, closeErr
	}

	return body, nil
}

func RequestBodyJsonToStruct(body []byte) interface{} {
	var jsonData interface{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		jsonData = ""
	}

	return jsonData
}

type httpRequest struct {
	Method     string              `json:"method"`
	Body       string              `json:"body"`
	ParsedBody interface{}         `json:"parsed_body"`
	Path       string              `json:"path"`
	Version    string              `json:"version"`
	Headers    map[string][]string `json:"headers"`
}

type opaInput struct {
	Input httpRequest `json:"input"`
}

func RequestToOpaInput(req *http.Request) ([]byte, error) {
	body, err := RequestBodyToString(req)
	if err != nil {
		return nil, err
	}

	parsedBody := RequestBodyJsonToStruct(body)

	path := ""
	if req.URL != nil {
		path = req.URL.Path
	}

	receivedRequest := httpRequest{
		Method:     req.Method,
		Body:       string(body),
		ParsedBody: parsedBody,
		Path:       path,
		Version:    fmt.Sprintf("%d.%d", req.ProtoMajor, req.ProtoMinor),
		Headers:    req.Header,
	}

	return json.Marshal(opaInput{receivedRequest})
}
