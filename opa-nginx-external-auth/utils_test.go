package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStringFromBody(t *testing.T) {
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
		body := ioutil.NopCloser(bytes.NewBuffer([]byte(c.bodyString)))

		reqBytes, err := GetStringFromBody(body)
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
