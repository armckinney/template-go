package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_HelloWorld(t *testing.T) {
	s := &Server{}
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	s.HelloWorldHandler(w, req)

	// Check StatusCode
	if w.Code != http.StatusOK {
		t.Errorf("got status %d; want %d", w.Code, http.StatusOK)
	}

	// Check Content-Type
	if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("got invalid content type %s; want application/json", contentType)
	}

	// Check Body
	var response map[string]string
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("cannot decode response: %v", err)
	}

	expectedMsg := "Hello World"
	if response["message"] != expectedMsg {
		t.Errorf("got message %s; want %s", response["message"], expectedMsg)
	}
}

func TestHandler_Health(t *testing.T) {
	s := &Server{}
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	s.healthHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("got status %d; want %d", w.Code, http.StatusOK)
	}

	var response map[string]string
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("cannot decode response: %v", err)
	}

	if response["status"] != "OK" {
		t.Errorf("got status %s; want %s", response["status"], "OK")
	}
}

func TestRegisterRoutes(t *testing.T) {
	s := &Server{}
	handler := s.RegisterRoutes()

	// Create a test server using the handler
	ts := httptest.NewServer(handler)
	defer ts.Close()

	// Test "/" route
	res, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("got status %d; want %d", res.StatusCode, http.StatusOK)
	}

	// Test "/health" route
	resHealth, err := http.Get(ts.URL + "/health")
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
	defer resHealth.Body.Close()

	if resHealth.StatusCode != http.StatusOK {
		t.Errorf("got status %d; want %d", resHealth.StatusCode, http.StatusOK)
	}
}
