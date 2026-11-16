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
		t.Fatalf("status esperado 200, veio %d", rec.Code)
	}
	var body helloResponse
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("resposta não é JSON: %v", err)
	}
	if body.Message != "hello world" {
		t.Errorf("esperado 'hello world', veio %q", body.Message)
	}
}

func TestHelloRejectsPost(t *testing.T) {
	rec := httptest.NewRecorder()
	newRouter().ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/hello", nil))

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status esperado 405, veio %d", rec.Code)
	}
}