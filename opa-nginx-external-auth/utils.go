package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jmespath/go-jmespath"
	"github.com/open-policy-agent/opa/bundle"
	"github.com/open-policy-agent/opa/rego"
)

func GetStringFromBody(input io.ReadCloser) (string, error) {
	if input == nil {
		return "", fmt.Errorf("Body is nil")
	}

	body, err := ioutil.ReadAll(input)
	if err != nil {
		return "", err
	}

	err = input.Close()
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

func NewHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
		},
		Timeout: 10 * time.Second,
	}
}

func NewJmsepathClient(expression string) (*jmespath.JMESPath, error) {
	return jmespath.Compile(expression)
}

//go:embed rego/*
var content embed.FS

type OpaClient struct {
	PartialResult rego.PartialResult
}

func NewOpaClient(ctx context.Context) (*OpaClient, error) {
	loader, err := bundle.NewFSLoader(content)
	if err != nil {
		return nil, err
	}

	reader := bundle.NewCustomReader(loader).WithSkipBundleVerification(true)
	b, err := reader.Read()
	if err != nil {
		return nil, err
	}

	r := rego.New(
		rego.ParsedBundle("bundle", &b),
		rego.Query(`data.nginx.authz`),
	)

	pr, err := r.PartialResult(ctx)
	if err != nil {
		return nil, err
	}

	return &OpaClient{
		PartialResult: pr,
	}, nil
}
