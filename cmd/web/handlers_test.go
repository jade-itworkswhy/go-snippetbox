package main

import (
	"jade-factory/go-snippetbox/internal/assert"
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	t.Parallel() // run test concurrently

	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}
