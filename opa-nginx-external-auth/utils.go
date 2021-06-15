package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jmespath/go-jmespath"
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
