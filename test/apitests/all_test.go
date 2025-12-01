package apitests

import (
	"fmt"
	"net/http"
	"testing"
)

var initCatId string

func init() {
	// Preparation: delete all existing & create a cat
	ids := []string{}
	call("GET", "/cats", nil, nil, &ids)

	for _, id := range ids {
		code := 0
		call("DELETE", "/cats/"+id, nil, &code, nil)
		fmt.Println("DELETE /cats ->", code)
	}

	// Create a single cat into the DB
	call("POST", "/cats", &CatModel{Name: "Toto"}, nil, &initCatId)
}

func TestGetCats(t *testing.T) {

	code := 0
	result := []string{}
	err := call("GET", "/cats", nil, &code, &result)
	if err != nil {
		t.Error("Request error", err)
	}

	fmt.Println("GET /cats ->", code, result)

	if code != http.StatusOK {
		t.Error("We should get code 200, got", code)
	}

	// On vérifie qu'il y a au moins 1 chat
	if len(result) < 1 {
		t.Error("Expected at least 1 item, got", len(result))
		return
	}

	// Vérifie que le chat créé initialement est présent
	found := false
	for _, id := range result {
		if id == initCatId {
			found = true
			break
		}
	}
	if !found {
		t.Error("initCatId not found in the list of cats")
	}
}

func TestCreateCat(t *testing.T) {
	code := 0
	var newID string
	cat := CatModel{Name: "Mimi"}

	err := call("POST", "/cats", &cat, &code, &newID)
	if err != nil {
		t.Error("Request error", err)
	}

	fmt.Println("POST /cats ->", code, newID)

	if code != http.StatusCreated {
		t.Error("Expected 201, got", code)
	}
	if newID == "" {
		t.Error("Expected new cat ID, got empty")
	}
}

func TestGetCatByID(t *testing.T) {
	code := 0
	var result CatModel

	err := call("GET", "/cats/"+initCatId, nil, &code, &result)
	if err != nil {
		t.Error("Request error", err)
	}

	fmt.Println("GET /cats/"+initCatId+" ->", code, result)

	if code != http.StatusOK {
		t.Error("Expected 200, got", code)
	}
	if result.ID != initCatId {
		t.Error("Expected ID", initCatId, "got", result.ID)
	}
}

func TestDeleteCat(t *testing.T) {
	code := 0
	err := call("DELETE", "/cats/"+initCatId, nil, &code, nil)
	if err != nil {
		t.Error("Request error", err)
	}

	fmt.Println("DELETE /cats/"+initCatId+" ->", code)

	if code != http.StatusOK && code != http.StatusNoContent {
		t.Error("Expected 200 or 204, got", code)
	}
}
