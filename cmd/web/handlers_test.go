package main

import (

	"net/http"
	"snippetbox/internal/assert"
	"testing"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	testStatusCode, _, testBody := ts.get(t, "/ping")
	assert.Equal(t, testStatusCode, http.StatusOK)
	assert.Equal(t, testBody, "OK")
}
