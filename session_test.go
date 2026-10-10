package gssion

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewCreatesSession(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	sess, err := New(res, req)
	if err != nil {
		t.Fatalf("New() returned error: %v", err)
	}
	if sess == nil {
		t.Fatal("New() returned nil session")
	}
	if sess.ID == "" {
		t.Fatal("session ID is empty")
	}
	if !sess.IsAuthenticated {
		// default should be false until login is called
	}
}

func TestGetReturnsSessionForValidCookie(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	sess, err := New(res, req)
	if err != nil {
		t.Fatalf("New() returned error: %v", err)
	}

	cookie := res.Result().Cookies()[0]
	req.AddCookie(cookie)

	loaded, err := Get(res, req)
	if err != nil {
		t.Fatalf("Get() returned error for valid session: %v", err)
	}
	if loaded == nil {
		t.Fatal("Get() returned nil for valid session")
	}
	if loaded.ID != sess.ID {
		t.Fatalf("expected session ID %q, got %q", sess.ID, loaded.ID)
	}
}

func TestGetFailsWithoutSessionCookie(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	if _, err := Get(res, req); err == nil {
		t.Fatal("Get() should fail without session cookie")
	}
}

func TestLoginSetsAuthenticationState(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	sess, err := New(res, req)
	if err != nil {
		t.Fatalf("New() returned error: %v", err)
	}

	sess.Login()
	if !sess.IsLoggedIn() {
		t.Fatal("Login() did not set IsAuthenticated")
	}
}
