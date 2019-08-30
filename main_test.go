package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEchoSuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(echoHandler))
	defer ts.Close()

	input := "{\"message\":\"Hello World!\"}"
	resp, err := http.Post(ts.URL, "application/json", strings.NewReader(input))
	if err != nil {
		t.Errorf("Failed to access the server")
	}

	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Failed to read the body")
	}

	if strings.TrimSpace(string(output)) != input {
		t.Errorf("output %s doesn't match input %s", output, input)
	}
}
