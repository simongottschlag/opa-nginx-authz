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

func TestGetStringFromRequestBody(t *testing.T) {
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

		reqBytes, err := GetStringFromRequestBody(testRequest)
		require.NoError(t, err)

		require.Equal(t, c.bodyString, string(reqBytes))
	}
}

func TestGetStructFromBody(t *testing.T) {
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
		structResponse := GetStructFromBody([]byte(c.bodyString))
		expectedJson, err := json.Marshal(c.expectedStruct)
		require.NoError(t, err)

		responseJson, err := json.Marshal(structResponse)
		require.NoError(t, err)

		require.JSONEq(t, string(expectedJson), string(responseJson))
	}
}

func TestGetOpaInputJson(t *testing.T) {
	cases := []struct {
		testDescription string
		bodyString      string
		request         *http.Request
		expectedStruct  OpaInput
	}{
		{
			testDescription: "empty",
			request:         &http.Request{},
			bodyString:      "",
			expectedStruct: OpaInput{
				Input{
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
			expectedStruct: OpaInput{
				Input{
					Method:     "GET",
					Body:       "",
					ParsedBody: "",
					Path:       "/foo",
					Version:    "1.1",
					Headers:    map[string][]string{"Foo": {"Bar"}},
				},
			},
		},
		{
			testDescription: "one parameter body",
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
			bodyString: `{"foo": "bar"}`,
			expectedStruct: OpaInput{
				Input{
					Method: "GET",
					Body:   `{"foo": "bar"}`,
					ParsedBody: map[string]interface{}{
						"foo": "bar",
					},
					Path:    "/foo",
					Version: "1.1",
					Headers: map[string][]string{"Foo": {"Bar"}},
				},
			},
		},
	}

	for i, c := range cases {
		t.Logf("Test iteration %d: %s", i, c.testDescription)
		testReq := c.request
		testReq.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(c.bodyString)))
		input, err := GetOpaInputJson(testReq)
		require.NoError(t, err)

		expectedJson, err := json.Marshal(c.expectedStruct)
		require.NoError(t, err)

		require.JSONEq(t, string(expectedJson), string(input))
	}
}

func TestGetOpaRequest(t *testing.T) {
	cases := []struct {
		testDescription string
		endpoint        string
	}{
		{
			testDescription: "empty endpoint",
			endpoint:        "",
		},
		{
			testDescription: "endpoint with port",
			endpoint:        "http://opa:1234/v1/data/test/abc",
		},
	}

	for i, c := range cases {
		t.Logf("Test iteration %d: %s", i, c.testDescription)
		testReq := &http.Request{
			Method: http.MethodGet,
			URL: &url.URL{
				Scheme: "http",
				Host:   "test.run",
				Path:   "/abc/123",
			},
			Header: http.Header{
				"Foo": {"Bar"},
			},
			ProtoMajor: 1,
			ProtoMinor: 1,
			Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(`{"foo": "bar"}`))),
		}

		opaReq, err := GetOpaRequest(testReq, c.endpoint)
		require.NoError(t, err)

		body, err := ioutil.ReadAll(opaReq.Body)
		require.NoError(t, err)

		err = opaReq.Body.Close()
		require.NoError(t, err)

		var opaReqInput OpaInput
		err = json.Unmarshal(body, &opaReqInput)
		require.NoError(t, err)

		require.Equal(t, "application/json", opaReq.Header.Get("Content-Type"))
		require.Equal(t, "POST", opaReq.Method)
		require.Equal(t, c.endpoint, opaReq.URL.String())
		require.Equal(t, `{"foo": "bar"}`, opaReqInput.Body)
		require.Equal(t, "GET", opaReqInput.Method)
		require.Equal(t, map[string]interface{}{"foo": "bar"}, opaReqInput.ParsedBody)
		require.Equal(t, "/abc/123", opaReqInput.Path)
		require.Equal(t, "1.1", opaReqInput.Version)
		require.Equal(t, map[string][]string{"Foo": {"Bar"}}, opaReqInput.Headers)
	}
}
