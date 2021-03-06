package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOpaProxyHandler(t *testing.T) {
	cases := []struct {
		testDescription     string
		authorizationHeader string
		expectedStatus      int
	}{
		{
			testDescription:     "status ok",
			authorizationHeader: "true",
			expectedStatus:      http.StatusOK,
		},
		{
			testDescription:     "status forbidden",
			authorizationHeader: "false",
			expectedStatus:      http.StatusForbidden,
		},
	}

	for i, c := range cases {
		t.Logf("Test iteration %d: %s", i, c.testDescription)
		testServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			body, err := GetStringFromBody(req.Body)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}

			var opaReqInput OpaInput
			err = json.Unmarshal([]byte(body), &opaReqInput)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}

			rw.Header().Add("Content-Type", "application/json")

			authz := opaReqInput.Headers["Authorization"][0]
			if authz != "true" {
				rw.WriteHeader(http.StatusForbidden)
				if _, err := rw.Write([]byte(`{"result": false}`)); err != nil {
					fmt.Fprintf(os.Stderr, "Could not write response data: %s", err)
				}

				return
			}

			if _, err := rw.Write([]byte(`{"result": true}`)); err != nil {
				fmt.Fprintf(os.Stderr, "Could not write response data: %s", err)
			}
			rw.WriteHeader(http.StatusOK)
		}))
		defer testServer.Close()

		httpClient := NewHttpClient()

		jmsepathClient, err := NewJmsepathClient("result")
		require.NoError(t, err)

		ctx := context.Background()
		opaClient, err := NewOpaClient(ctx)
		require.NoError(t, err)

		handlerClient := NewHandlerClient(httpClient, jmsepathClient, opaClient, testServer.URL)

		req, err := http.NewRequest("GET", "/", nil)
		require.NoError(t, err)

		req.Header.Add("Authorization", c.authorizationHeader)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlerClient.OpaProxyHandler)

		handler.ServeHTTP(rr, req)

		require.Equal(t, c.expectedStatus, rr.Code)
	}
}

func TestOpaRegoHandler(t *testing.T) {
	cases := []struct {
		testDescription     string
		authorizationHeader string
		expectedStatus      int
	}{
		{
			testDescription:     "status ok",
			authorizationHeader: "Bearer test",
			expectedStatus:      http.StatusOK,
		},
		{
			testDescription:     "status forbidden",
			authorizationHeader: "Bearer abc",
			expectedStatus:      http.StatusForbidden,
		},
	}

	for i, c := range cases {
		t.Logf("Test iteration %d: %s", i, c.testDescription)
		httpClient := NewHttpClient()

		jmsepathClient, err := NewJmsepathClient("result")
		require.NoError(t, err)

		ctx := context.Background()
		opaClient, err := NewOpaClient(ctx)
		require.NoError(t, err)

		handlerClient := NewHandlerClient(httpClient, jmsepathClient, opaClient, "http://foo.bar/baz")

		req, err := http.NewRequest("GET", "/", nil)
		require.NoError(t, err)

		req.Header.Add("Authorization", c.authorizationHeader)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlerClient.OpaRegoHandler)

		handler.ServeHTTP(rr, req)

		require.Equal(t, c.expectedStatus, rr.Code)
	}
}
