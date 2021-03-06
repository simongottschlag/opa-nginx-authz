package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	testBody    = `{"foo": "bar"}`
	testHeaders = map[string][]string{"Foo": {"Bar"}}
	testReq     = &http.Request{
		Method: http.MethodGet,
		URL: &url.URL{
			Scheme: "http",
			Host:   "test.run",
			Path:   "/abc/123",
		},
		Header:     testHeaders,
		ProtoMajor: 1,
		ProtoMinor: 1,
	}
)

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
		req := c.request
		req.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(c.bodyString)))
		input, err := GetOpaInputJson(req)
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

		req := testReq
		req.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(testBody)))

		opaReq, err := GetOpaRequest(req, c.endpoint)
		require.NoError(t, err)

		body, err := GetStringFromBody(opaReq.Body)
		require.NoError(t, err)

		var opaReqInput OpaInput
		err = json.Unmarshal([]byte(body), &opaReqInput)
		require.NoError(t, err)

		require.Equal(t, "application/json", opaReq.Header.Get("Content-Type"))
		require.Equal(t, "POST", opaReq.Method)
		require.Equal(t, c.endpoint, opaReq.URL.String())
		require.Equal(t, testBody, opaReqInput.Body)
		require.Equal(t, testReq.Method, opaReqInput.Method)
		require.Equal(t, map[string]interface{}{"foo": "bar"}, opaReqInput.ParsedBody)
		require.Equal(t, testReq.URL.Path, opaReqInput.Path)
		require.Equal(t, fmt.Sprintf("%d.%d", testReq.ProtoMajor, testReq.ProtoMinor), opaReqInput.Version)
		require.Equal(t, testHeaders, opaReqInput.Headers)
	}
}

func TestGetOpaResponse(t *testing.T) {
	cases := []struct {
		testDescription        string
		testPath               string
		testServerResponseCode int
	}{
		{
			testDescription:        "empty path, status ok",
			testPath:               "",
			testServerResponseCode: http.StatusOK,
		},
		{
			testDescription:        "empty path, status forbidden",
			testPath:               "",
			testServerResponseCode: http.StatusForbidden,
		},
		{
			testDescription:        "non-empty path, status ok",
			testPath:               "/abc/123",
			testServerResponseCode: http.StatusOK,
		},
		{
			testDescription:        "non-empty path, status forbidden",
			testPath:               "/abc/123",
			testServerResponseCode: http.StatusForbidden,
		},
	}

	for i, c := range cases {
		t.Logf("Test iteration %d: %s", i, c.testDescription)
		testServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(c.testServerResponseCode)
		}))
		defer testServer.Close()

		httpClient := NewHttpClient()

		endpoint := fmt.Sprintf("%s/%s", testServer.URL, c.testPath)

		req := testReq
		req.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(testBody)))

		opaResponse, err := GetOpaHttpResponse(httpClient, req, endpoint)
		require.NoError(t, err)

		require.Equal(t, c.testServerResponseCode, opaResponse.StatusCode)
		require.Equal(t, fmt.Sprintf("/%s", c.testPath), opaResponse.Request.URL.Path)
	}
}

func TestGetOpaResponseStruct(t *testing.T) {
	cases := []struct {
		testDescription string
		opaResponse     string
		expectedResult  bool
	}{
		{
			testDescription: "result true",
			opaResponse:     `{"result": true}`,
			expectedResult:  true,
		},
		{
			testDescription: "result false",
			opaResponse:     `{"result": false}`,
			expectedResult:  false,
		},
	}

	for i, c := range cases {
		t.Logf("Test iteration %d: %s", i, c.testDescription)
		testServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusOK)
			if _, err := rw.Write([]byte(c.opaResponse)); err != nil {
				fmt.Fprintf(os.Stderr, "Could not write response data: %s", err)
			}
		}))
		defer testServer.Close()

		httpClient := NewHttpClient()

		req := testReq
		req.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(testBody)))

		responseStruct, err := GetOpaResponseStruct(httpClient, req, testServer.URL)
		require.NoError(t, err)

		parsedStruct, ok := responseStruct.(map[string]interface{})
		require.Equal(t, true, ok)

		result, ok := parsedStruct["result"].(bool)
		require.Equal(t, true, ok)

		require.Equal(t, c.expectedResult, result)
	}
}

func TestGetResultFromOpaResponseStruct(t *testing.T) {
	cases := []struct {
		testDescription string
		opaResponse     string
		expectedResult  bool
	}{
		{
			testDescription: "result true",
			opaResponse:     `{"result": true}`,
			expectedResult:  true,
		},
		{
			testDescription: "result false",
			opaResponse:     `{"result": false}`,
			expectedResult:  false,
		},
	}

	for i, c := range cases {
		t.Logf("Test iteration %d: %s", i, c.testDescription)
		testServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusOK)
			if _, err := rw.Write([]byte(c.opaResponse)); err != nil {
				fmt.Fprintf(os.Stderr, "Could not write response data: %s", err)
			}
		}))
		defer testServer.Close()

		httpClient := NewHttpClient()

		jmsepathClient, err := NewJmsepathClient("result")
		require.NoError(t, err)

		req := testReq
		req.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(testBody)))

		responseStruct, err := GetOpaResponseStruct(httpClient, req, testServer.URL)
		require.NoError(t, err)

		result, err := GetResultFromOpaResponseStruct(jmsepathClient, responseStruct)
		require.NoError(t, err)

		require.Equal(t, c.expectedResult, result)
	}
}

func BenchmarkGetResultFromOpaResponseStructMock(b *testing.B) {
	testServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		if _, err := rw.Write([]byte(`{"result": true}`)); err != nil {
			fmt.Fprintf(os.Stderr, "Could not write response data: %s", err)
		}
	}))
	defer testServer.Close()

	httpClient := NewHttpClient()

	jmsepathClient, err := NewJmsepathClient("result")
	require.NoError(b, err)

	req := testReq
	req.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(testBody)))

	responseStruct, err := GetOpaResponseStruct(httpClient, req, testServer.URL)
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		result, err := GetResultFromOpaResponseStruct(jmsepathClient, responseStruct)
		require.NoError(b, err)
		require.Equal(b, true, result)
	}
}

func BenchmarkGetResultFromOpaResponseStruct(b *testing.B) {
	httpClient := NewHttpClient()

	jmsepathClient, err := NewJmsepathClient("result")
	require.NoError(b, err)

	req := testReq
	req.Header.Add("Authorization", "Bearer test")
	req.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(testBody)))

	_, err = GetOpaResponseStruct(httpClient, req, "http://localhost:8181/v1/data/nginx/authz")
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		responseStruct, err := GetOpaResponseStruct(httpClient, req, "http://localhost:8181/v1/data/nginx/authz")
		require.NoError(b, err)

		result, err := GetResultFromOpaResponseStruct(jmsepathClient, responseStruct)
		require.NoError(b, err)
		require.Equal(b, true, result)
	}
}

func TestGetResultWithOpaInput(t *testing.T) {
	cases := []struct {
		testDescription string
		opaInput        Input
		expectedResult  bool
	}{
		{
			testDescription: "result true",
			opaInput: Input{
				Method:     "GET",
				Body:       "",
				ParsedBody: "",
				Path:       "/",
				Version:    "",
				Headers:    map[string][]string{"Authorization": {"Bearer test"}},
			},
			expectedResult: true,
		},
		{
			testDescription: "result false",
			opaInput: Input{
				Method:     "GET",
				Body:       "",
				ParsedBody: "",
				Path:       "/",
				Version:    "",
				Headers:    map[string][]string{"Authorization": {"Bearer false"}},
			},
			expectedResult: false,
		},
	}

	for i, c := range cases {
		t.Logf("Test iteration %d: %s", i, c.testDescription)
		ctx := context.Background()
		opaClient, err := NewOpaClient(ctx)
		require.NoError(t, err)

		result, err := GetResultWithOpaInput(ctx, opaClient, c.opaInput)
		require.NoError(t, err)

		require.Equal(t, c.expectedResult, result)
	}
}

func BenchmarkGetResultWithOpaInput(b *testing.B) {
	input := Input{
		Method:     "GET",
		Body:       "",
		ParsedBody: "",
		Path:       "/",
		Version:    "",
		Headers:    map[string][]string{"Authorization": {"Bearer test"}},
	}

	ctx := context.Background()
	opaClient, err := NewOpaClient(ctx)
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		result, err := GetResultWithOpaInput(ctx, opaClient, input)
		require.NoError(b, err)
		require.Equal(b, true, result)
	}
}
