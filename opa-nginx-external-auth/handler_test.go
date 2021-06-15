package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOpaHandler(t *testing.T) {
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
				rw.Write([]byte(`{"result": false}`))
				return
			}

			rw.Write([]byte(`{"result": true}`))
			rw.WriteHeader(http.StatusOK)
		}))
		defer testServer.Close()

		httpClient := NewHttpClient()

		jmsepathClient, err := NewJmsepathClient("result")
		require.NoError(t, err)

		handlerClient := NewHandlerClient(httpClient, jmsepathClient, testServer.URL)

		req, err := http.NewRequest("GET", "/", nil)
		require.NoError(t, err)

		req.Header.Add("Authorization", c.authorizationHeader)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlerClient.OpaHandler)

		handler.ServeHTTP(rr, req)

		require.Equal(t, c.expectedStatus, rr.Code)
	}
}
