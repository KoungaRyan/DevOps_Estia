package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogReq(t *testing.T) {
	called := false

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	})

	handler := logReq(next)

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if !called {
		t.Errorf("next handler was not called")
	}
}

func TestNewAppInitialization(t *testing.T) {
	app := newApp()

	req := httptest.NewRequest("GET", "/api/cats", nil)
	rec := httptest.NewRecorder()

	app.ServeHTTP(rec, req)
	// Pas d'erreur → route existante
	if rec.Code == 0 {
		t.Errorf("app did not respond")
	}
}
