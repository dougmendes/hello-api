package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	rec := httptest.NewRecorder()
	newRouter().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/hello", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	var body helloResponse
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("response is not valid JSON: %v", err)
	}
	if body.Message != "hello world" {
		t.Errorf("expected 'hello world', got %q", body.Message)
	}
}

func TestHelloRejectsPost(t *testing.T) {
	rec := httptest.NewRecorder()
	newRouter().ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/hello", nil))

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405, got %d", rec.Code)
	}
}