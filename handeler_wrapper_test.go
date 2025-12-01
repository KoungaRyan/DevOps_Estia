package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMakeHandlerFunc_Happy(t *testing.T) {
	svc := func(r *http.Request) (int, any) {
		return http.StatusOK, Cat{ID: "1", Name: "Tom"}
	}

	handler := makeHandlerFunc(svc)

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}

	var c Cat
	json.NewDecoder(rec.Body).Decode(&c)

	if c.Name != "Tom" {
		t.Errorf("unexpected body: %v", c)
	}
}

func TestMakeHandlerFunc_Panic(t *testing.T) {
	svc := func(r *http.Request) (int, any) {
		panic("boom")
	}

	handler := makeHandlerFunc(svc)

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", rec.Code)
	}
}
