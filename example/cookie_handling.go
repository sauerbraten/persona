package main

import (
	"net/http"
	"time"
)

// looks for a session cookie and returns the user's email address, or "" if something failed.
// this example has no error handling!
func getUserEmail(req *http.Request) (email string) {
	cookie, err := req.Cookie("session")
	if err != nil {
		return
	}

	err = secCookie.Decode("session", cookie.Value, &email)
	return
}

// sets a secure session cookie containing the user's email address, so we can recognize him later
func setSessionCookie(resp http.ResponseWriter, email string, expires int64) error {
	encoded, err := secCookie.Encode("session", email)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     "session",
		Value:    encoded,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 336),
	}

	http.SetCookie(resp, cookie)

	// add new (first-time) users to list of known users
	if !userExists(email) {
		addUser(email)
	}

	return nil
}

// overwrites the secure session cookie
func revokeSessionCookie(resp http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	}

	http.SetCookie(resp, cookie)
}
