package main

import (
	"net/http"
	"path"
	"strings"
)

// extractCatID attempts to get the cat ID from the request URL path.
// It falls back to taking the last path segment so tests that set
// the request URL (e.g. "/api/cats/1") continue to work.
func extractCatID(req *http.Request) string {
	// If the server or router provides a path value API (Go 1.25+), the tests
	// may also set it. However to keep compatibility, simply parse the URL.
	p := req.URL.Path
	// Ensure we don't return empty when the path ends with a slash
	p = strings.TrimSuffix(p, "/")
	return path.Base(p)
}

func getCat(req *http.Request) (int, any) {
	catId := extractCatID(req)
	Logger.Info("Getting cat", catId)

	if cat, found := catsDatabase[catId]; found {
		Logger.Info("Cat found")
		return http.StatusOK, cat
	}

	Logger.Info("Cat not found")
	return http.StatusNotFound, "Cat not found"
}

func deleteCat(req *http.Request) (int, any) {
	catId := extractCatID(req)
	Logger.Info("Deleting cat", catId)

	if _, found := catsDatabase[catId]; found {
		delete(catsDatabase, catId)
		Logger.Info("Cat deleted")
		return http.StatusNoContent, catId
	}

	Logger.Info("Cat not found")
	return http.StatusNotFound, "Cat not found"
}