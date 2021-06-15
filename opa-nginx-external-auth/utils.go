package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetStringFromRequestBody(req *http.Request) (string, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}

	err = req.Body.Close()
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetStructFromBody(body []byte) interface{} {
	var jsonData interface{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		jsonData = ""
	}

	return jsonData
}

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
	body, err := GetStringFromRequestBody(req)
	if err != nil {
		return nil, err
	}

	parsedBody := GetStructFromBody([]byte(body))

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
