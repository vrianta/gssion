# Gssion Session Management Runbook

This guide shows how to create a session, store session data, and access it in later requests using the `gssion` package.

## 1. Install

```bash
go get github.com/vrianta/gssion
```

Then import it in your Go application:

```go
import "github.com/vrianta/gssion"
```

## 2. What this library does

The package creates a cookie called `sessionid` and stores session data in memory. Each request can read the current session from that cookie using `gssion.Get(...)`.

Session data is stored in the `Context.Data` map:

```go
session.Data["user_id"] = 42
session.Data["name"] = "Alice"
```

## 3. Create a new session

Use `gssion.New(w, r)` in a handler when the user first arrives or logs in.

```go
package main

import (
    "fmt"
    "net/http"

    gssion "github.com/vrianta/gssion"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
    session, err := gssion.New(w, r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    session.Data["user_id"] = 101
    session.Data["name"] = "Alice"
    session.Login()

    fmt.Fprintln(w, "logged in")
}
```

Notes:

- `gssion.New(...)` creates a new session ID and sets the `sessionid` cookie automatically.
- `session.Login()` marks the session as authenticated.
- `session.Data` is the place to store values like user ID, role, cart data, and preferences.

## 4. Access an existing session

On later requests, read the cookie from the incoming request and load the same session:

```go
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
    session, err := gssion.Get(w, r)
    if err != nil {
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }

    if !session.IsLoggedIn() {
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }

    userID, _ := session.Data["user_id"].(int)
    name, _ := session.Data["name"].(string)

    fmt.Fprintf(w, "Welcome %s (user %d)", name, userID)
}
```

If the session cookie is missing, `gssion.Get(...)` will fail, and you can redirect the user back to login.

## 5. Read the session ID directly

If you need the raw session ID:

```go
func checkSessionID(r *http.Request) {
    sessionID, err := gssion.GetSessionID(r)
    if err != nil {
        fmt.Println("No session ID found")
        return
    }

    fmt.Println("Session ID:", sessionID)
}
```

## 6. Store and read data

```go
session, _ := gssion.New(w, r)

session.Data["email"] = "alice@example.com"
session.Data["role"] = "admin"

email, _ := session.Data["email"].(string)
role, _ := session.Data["role"].(string)

fmt.Println(email)
fmt.Println(role)
```

## 7. Logout and remove session

```go
func logoutHandler(w http.ResponseWriter, r *http.Request) {
    session, err := gssion.Get(w, r)
    if err == nil {
        session.RemoveSession()
    }

    http.Redirect(w, r, "/login", http.StatusFound)
}
```

You can also clear the browser cookie manually:

```go
session, _ := gssion.Get(w, r)
session.Logout()
```

## 8. Typical full example

```go
package main

import (
    "fmt"
    "net/http"

    gssion "github.com/vrianta/gssion"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
    session, err := gssion.New(w, r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    session.Data["user_id"] = 7
    session.Data["username"] = "demo-user"
    session.Login()

    fmt.Fprintln(w, "Login successful")
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
    session, err := gssion.Get(w, r)
    if err != nil {
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }

    if !session.IsLoggedIn() {
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }

    username, _ := session.Data["username"].(string)
    fmt.Fprintf(w, "Hello %s", username)
}

func main() {
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/profile", profileHandler)
    http.ListenAndServe(":8080", nil)
}
```

## 9. Best practices

- Use `gssion.New(...)` when the user logs in or starts a new session.
- Use `gssion.Get(...)` on every authenticated route.
- Keep user-sensitive values in `session.Data` instead of storing them in global variables.
- Use HTTPS in production if you enable secure cookies.
- Call `RemoveSession()` or `Logout()` when the user signs out.

## 10. Important API summary

- `gssion.New(w, r)` — create a new session
- `gssion.Get(w, r)` — load the current session from the request cookie
- `gssion.GetSessionID(r)` — read only the session ID
- `session.Data[...]` — store custom session values
- `session.Login()` — mark the session as authenticated
- `session.IsLoggedIn()` — check authentication state
- `session.RemoveSession()` — delete the session
- `session.Logout()` — clear the cookie

This is the standard flow:

1. User hits a login page
2. Create session with `gssion.New(...)`
3. Store user info in `session.Data`
4. Call `session.Login()`
5. On next request, call `gssion.Get(...)`
6. Read data from `session.Data`
7. Logout or remove the session when done
