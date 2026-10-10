# Changelog

## v1.0.0 - 2026-10-10

### Overview
This is the first public release of `gssion`, a lightweight Go session management library for HTTP handlers.

### Included features
- Session creation with `gssion.New(w, r)`
- Session retrieval with `gssion.Get(w, r)`
- Session ID access via `gssion.GetSessionID(r)`
- Custom session data storage in `session.Data`
- Login and logout helpers via `session.Login()` and `session.Logout()`
- Cookie-based session tracking using the `sessionid` cookie
- Session removal and cleanup support
- In-memory session storage with session expiry handling
- Session persistence to a local disk file (`sessions.data`)

### Behavior
- Each new session receives a unique random ID.
- Session data is stored per request context and can be accessed across requests through the session cookie.
- A session can be marked authenticated using `Login()`.
- Expired sessions are cleaned up automatically by the package runtime.

### Notes
- This version is intended for lightweight server-side session handling in Go web applications.
- It is suitable for applications that need simple per-user state without a full external session store.
- For production deployments, consider reviewing the storage strategy and cookie security settings.

### Initial release highlights
- Easy API for creating and reading sessions
- Request-scoped access patterns for HTTP handlers
- Basic session lifecycle management for authentication flows
- Built-in cleanup and persistence support
