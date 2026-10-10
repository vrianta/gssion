package gssion

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"
)

func GetCookie(cookie_name string, r *http.Request) (*http.Cookie, error) {
	if cookie, err := r.Cookie(cookie_name); err == nil {
		// Log("Cookie: ", cookie)
		return cookie, nil
	} else {
		return nil, nil
	}
}

// Return true if it created te cookie else false means cookie is already present in the session
func AddCookie(cookie_config *http.Cookie, w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, cookie_config)
	// cookie_header := formHeader(cookie_config)
	// w.Header().Add("Set-Cookie", cookie_header)
	// w.Header().Add("X-Custom-Header", "MyHeaderValue")
	// w.Header().Set("Set-Cookie", cookie_header)
}

func RemoveCookie(cookie_name string, w http.ResponseWriter, r *http.Request) {
	cookie_header := fmt.Sprintf("%s=expire; Max-Age=-1; Expires=%s;", cookie_name, time.Now().UTC().Format(http.TimeFormat))
	if _, err := r.Cookie(cookie_name); err == nil {
		w.Header().Add("Set-Cookie", cookie_header)
	}
}

func formHeader(cookie_config *http.Cookie) string {

	cookie_header := fmt.Sprintf(
		"%s=%s;",
		cookie_config.Name,
		cookie_config.Value,
	)

	if !cookie_config.Expires.IsZero() {
		cookie_header += fmt.Sprintf(" Expires=%s;", cookie_config.Expires.Format(http.TimeFormat))
	} else if cookie_config.MaxAge == -1 {
		cookie_header += fmt.Sprintf(" Max-Age=%d;", 0)
	}
	if cookie_config.HttpOnly {
		cookie_header += " HttpOnly;"
	}
	if cookie_config.Secure {
		cookie_header += " Secure;"
	}

	return cookie_header
}

// Sets the session cookie in the client's browser
func (sh *Context) setCookie(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	sh.ExpirationTime = now.Add(30 * time.Minute).UTC()
	c := &http.Cookie{
		Name:     "sessionid",
		Value:    sh.ID,
		HttpOnly: true,
		Secure:   config.Https,
		SameSite: http.SameSiteLaxMode,
		Expires:  sh.ExpirationTime,
		Path:     "/",
	}

	AddCookie(c, w, r)
}

// GenerateSessionID generates a random session ID
func generateSessionID() (string, error) {
	// Create a byte slice to hold the random data
	bytes := make([]byte, 16) // 16 bytes = 128 bits, reasonable for session ID

	// Generate random bytes using crypto/rand
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a hexadecimal string
	sessionID := hex.EncodeToString(bytes)
	return sessionID, nil
}

// GenerateSessionID generates a random session ID
// func generateSessionID() (string, error) {
// 	// Create a byte slice to hold the random data
// 	bytes := make([]byte, 16) // 16 bytes = 128 bits, reasonable for session ID

// 	// Generate random bytes using crypto/rand
// 	_, err := rand.Read(bytes)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Convert the byte slice to a hexadecimal string
// 	sessionID := hex.EncodeToString(bytes)
// 	return sessionID, nil
// }
