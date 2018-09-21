package main

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"path"
	"strings"
	"time"
)

const (
	userEmail    = "me@example.com"
	userPassword = "xxx"
	cookieName   = "email"
	cookieExpiry = time.Minute * 60 * 24
)

// Note: In production I'd create a session to store the user's UUID in it.
func signinHandler(w http.ResponseWriter, r *http.Request) {
	var (
		email    = r.FormValue("email")
		password = r.FormValue("password")
	)

	// unauthenticated
	if email != userEmail || password != userPassword {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	// authenticated
	http.SetCookie(w, &http.Cookie{Name: cookieName, Value: email, Expires: time.Now().Add(cookieExpiry)})
	http.Redirect(w, r, "/private/main.html", http.StatusFound)
}

// Note: In production I'd get the UUID from the session and check if there's a corresponding user in the database.
func authHandler(w http.ResponseWriter, r *http.Request) {
	// get cookie
	if _, err := r.Cookie(cookieName); err != nil {
		http.Error(w, fmt.Sprintf("%d %s", http.StatusNotFound, http.StatusText(http.StatusNotFound)), http.StatusNotFound)
		return
	}

	// set Content-Type header
	if mime := mime.TypeByExtension(path.Ext(r.URL.Path)); mime != "" {
		w.Header().Set("Content-Type", mime)
	}

	// set X-Accel-Redirect header
	p := strings.Replace(r.URL.Path, "/private", "/internal", -1)
	w.Header().Set("X-Accel-Redirect", p)

	fmt.Printf("%s -> %s\n", r.URL.Path, p)
}

func main() {
	http.HandleFunc("/signin", signinHandler)
	http.HandleFunc("/", authHandler)

	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}
