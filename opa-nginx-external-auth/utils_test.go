package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRequestBodyToString(t *testing.T) {
	cases := []struct {
		testDescription string
		requestMethod   string
		requestEndpoint string
		bodyString      string
	}{
		{
			testDescription: "single word get",
			requestMethod:   http.MethodPost,
			requestEndpoint: "http://localhost",
			bodyString:      "test",
		},
		{
			testDescription: "multi word get",
			requestMethod:   http.MethodPost,
			requestEndpoint: "http://localhost",
			bodyString:      "test abc 123",
		},
		{
			testDescription: "multi word line get",
			requestMethod:   http.MethodPost,
			requestEndpoint: "http://localhost",
			bodyString: `test line one
			test line two
			test line three`,
		},
	}

	for i, c := range cases {
		t.Logf("Test iteration %d: %s", i, c.testDescription)
		testRequest, err := http.NewRequest(c.requestMethod, c.requestEndpoint, bytes.NewBuffer([]byte(c.bodyString)))
		require.NoError(t, err)

		reqBytes, err := RequestBodyToString(testRequest)
		require.NoError(t, err)

		require.Equal(t, c.bodyString, string(reqBytes))
	}
}

func TestRequestBodyJsonToStruct(t *testing.T) {
	cases := []struct {
		testDescription string
		bodyString      string
		expectedStruct  interface{}
	}{
		{
			testDescription: "empty body",
			bodyString:      "",
			expectedStruct:  "",
		},
		{
			testDescription: "single string parameter",
			bodyString:      `{"foo": "bar"}`,
			expectedStruct: map[string]interface{}{
				"foo": "bar",
			},
		},
		{
			testDescription: "single bool parameter",
			bodyString:      `{"foo": true}`,
			expectedStruct: map[string]interface{}{
				"foo": true,
			},
		},
		{
			testDescription: "two parameters",
			bodyString:      `{"foo": "bar", "baz": true}`,
			expectedStruct: map[string]interface{}{
				"foo": "bar",
				"baz": true,
			},
		},
		{
			testDescription: "array with one item",
			bodyString:      `{"foo": ["bar"]}`,
			expectedStruct: map[string]interface{}{
				"foo": []interface{}{"bar"},
			},
		},
		{
			testDescription: "array with two items",
			bodyString:      `{"foo": ["bar", "baz"]}`,
			expectedStruct: map[string]interface{}{
				"foo": []interface{}{"bar", "baz"},
			},
		},
		{
			testDescription: "object with one parameter",
			bodyString:      `{"foo": {"bar": "baz"}}`,
			expectedStruct: map[string]interface{}{
				"foo": map[string]interface{}{"bar": "baz"},
			},
		},
		{
			testDescription: "object with two parameters",
			bodyString:      `{"foo": {"bar": "baz", "foo": "bar"}}`,
			expectedStruct: map[string]interface{}{
				"foo": map[string]interface{}{"bar": "baz", "foo": "bar"},
			},
		},
	}

	for i, c := range cases {
		t.Logf("Test iteration %d: %s", i, c.testDescription)
		structResponse := RequestBodyJsonToStruct([]byte(c.bodyString))
		expectedJson, err := json.Marshal(c.expectedStruct)
		require.NoError(t, err)

		responseJson, err := json.Marshal(structResponse)
		require.NoError(t, err)

		require.JSONEq(t, string(expectedJson), string(responseJson))
	}
}

func TestRequestToOpaInput(t *testing.T) {
	cases := []struct {
		testDescription string
		bodyString      string
		request         *http.Request
		expectedStruct  opaInput
	}{
		{
			testDescription: "empty",
			request:         &http.Request{},
			bodyString:      "",
			expectedStruct: opaInput{
				httpRequest{
					Method:     "",
					Body:       "",
					ParsedBody: "",
					Path:       "",
					Version:    "0.0",
					Headers:    nil,
				},
			},
		},
		{
			testDescription: "empty body",
			request: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					Path: "/foo",
				},
				Header: http.Header{
					"Foo": {"Bar"},
				},
				ProtoMajor: 1,
				ProtoMinor: 1,
			},
			bodyString: "",
			expectedStruct: opaInput{
				httpRequest{
					Method:     "GET",
					Body:       "",
					ParsedBody: "",
					Path:       "/foo",
					Version:    "1.1",
					Headers:    map[string][]string{"Foo": {"Bar"}},
				},
			},
		},
	}

	for i, c := range cases {
		t.Logf("Test iteration %d: %s", i, c.testDescription)
		testReq := c.request
		testReq.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(c.bodyString)))
		input, err := RequestToOpaInput(testReq)
		require.NoError(t, err)

		expectedJson, err := json.Marshal(c.expectedStruct)
		require.NoError(t, err)

		require.JSONEq(t, string(expectedJson), string(input))
	}
}
