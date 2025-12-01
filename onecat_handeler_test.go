package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCat_Found(t *testing.T) {
	catsDatabase = map[string]Cat{
		"1": {ID: "1", Name: "Tom", Color: "Gray"},
	}

	req := httptest.NewRequest("GET", "/api/cats/1", nil)
	req.SetPathValue("catId", "1")

	code, body := getCat(req)

	if code != http.StatusOK {
		t.Errorf("expected 200, got %d", code)
	}
	cat, ok := body.(Cat)
	if !ok {
		t.Fatalf("expected body of type Cat, got %T", body)
	}
	if cat.Name != "Tom" {
		t.Errorf("expected Name=Tom, got %s", cat.Name)
	}
}

func TestGetCat_NotFound(t *testing.T) {
	catsDatabase = map[string]Cat{}

	req := httptest.NewRequest("GET", "/api/cats/404", nil)
	req.SetPathValue("catId", "404")

	code, body := getCat(req)

	if code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", code)
	}
	if body != "Cat not found" {
		t.Errorf("unexpected body: %v", body)
	}
}

func TestDeleteCat_Found(t *testing.T) {
	catsDatabase = map[string]Cat{
		"10": {ID: "10",
		 Name: "Garfield"},
	}

	req := httptest.NewRequest("DELETE", "/api/cats/10", nil)
	req.SetPathValue("catId", "10")

	code, body := deleteCat(req)

	if code != http.StatusNoContent {
		t.Errorf("expected 204, got %d", code)
	}
	if body != "10" {
		t.Errorf("expected 10, got %v", body)
	}
	if _, exists := catsDatabase["10"]; exists {
		t.Errorf("cat not deleted")
	}
}

func TestDeleteCat_NotFound(t *testing.T) {
	catsDatabase = map[string]Cat{}

	req := httptest.NewRequest("DELETE", "/api/cats/99", nil)
	req.SetPathValue("catId", "99")

	code, _ := deleteCat(req)

	if code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", code)
	}
}
