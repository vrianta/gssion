package gssion

import (
	"net/http"
	"time"
)

type (

	// HTTPPostParameters represents POST request parameters
	// hTTPPostParameters map[string]string

	// HTTPGetParameters represents GET request parameters
	// hTTPGetParameters map[string]string
	Context struct {
		ID string

		Data sessionData

		IsAuthenticated bool

		ExpirationTime time.Time

		w http.ResponseWriter
		r *http.Request
	}

	// Collection represents a collection of sessions managed as a heap
	collection []*Context

	// SessionData stores arbitrary key-value data for a session
	sessionData map[string]any

	// LRUCacheOperation represents an operation to update the LRU cache
	lRUCacheOperation struct {
		ID            string
		OperationType string // "move_to_front" or "insert" or "remove"
	}
)

type (
// viewT = func() view.Context
)
