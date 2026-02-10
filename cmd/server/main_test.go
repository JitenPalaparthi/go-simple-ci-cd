package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRoot(t *testing.T) {
	h := routes()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); !strings.Contains(ct, "text/plain") {
		t.Fatalf("expected text/plain content-type, got %q", ct)
	}
	if !strings.Contains(rec.Body.String(), "Hello from Go") {
		t.Fatalf("unexpected body: %q", rec.Body.String())
	}
}

func TestHealth(t *testing.T) {
	h := routes()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); !strings.Contains(ct, "application/json") {
		t.Fatalf("expected application/json content-type, got %q", ct)
	}

	var out healthResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &out); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if out.Status != "ok" {
		t.Fatalf("expected status ok, got %q", out.Status)
	}
	if out.Timestamp == "" {
		t.Fatalf("expected non-empty timestamp")
	}
}

func TestGetenv(t *testing.T) {
	t.Setenv("PORT", "9999")
	if got := getenv("PORT", "8080"); got != "9999" {
		t.Fatalf("expected 9999, got %s", got)
	}
	if got := getenv("MISSING", "x"); got != "x" {
		t.Fatalf("expected default, got %s", got)
	}
}
