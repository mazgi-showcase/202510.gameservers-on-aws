package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	handler(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Errorf("unexpected status code: got (%v) want (%v)", recorder.Code, http.StatusOK)
	}
}
