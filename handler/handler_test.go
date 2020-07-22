package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestHome(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Home)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("failed when hitting home. Status got %d want %d", status, http.StatusOK)
	}

	expected := "Welcome!\n"
	if rr.Body.String() != expected {
		t.Errorf("Failed to return expected response. Got %s want %s", rr.Body.String(), expected)
	}
}

func TestWelcome(t *testing.T) {
	path := "/welcome?name=Eph"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/welcome", Welcome).Queries("name", "{name}")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("failed when hitting welcome. Status got %d want %d", status, http.StatusOK)
	}

	expected := fmt.Sprintf("Welcome Eph!\nThis is a simple api webserver. Don't hope for too much!\n")
	if rr.Body.String() != expected {
		t.Errorf("Failed to return expected response. Got %s want %s", rr.Body.String(), expected)
	}
}

func TestHistory(t *testing.T) {
	req, err := http.NewRequest("GET", "/history", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(History)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("failed when hitting history. Status got %d want %d", status, http.StatusOK)
	}
}
