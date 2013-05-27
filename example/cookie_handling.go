package main

import (
	"github.com/gorilla/securecookie"
	"net/http"
	"time"
)

// initialize secure cookie storage; 16 byte ~ 128 bit AES encryption
var secCookie *securecookie.SecureCookie = securecookie.New([]byte("my very secret hash key"), []byte{0x5d, 0xa5, 0xd3, 0x90, 0xc9, 0x54, 0xa1, 0xc3, 0x70, 0x00, 0x8d, 0x6d, 0xa9, 0xd1, 0x07, 0x53})

// looks for a session cookie and returns the user's email address, or "" if something failed.
// this example has no error handling!
func getEmail(req *http.Request) (email string) {
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

	http.SetCookie(resp, &http.Cookie{
		Name:     "session",
		Value:    encoded,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 336), // session is valid for 2 weeks
	})

	// add new (first-time) users to list of known users
	if !userExists(email) {
		addUser(email)
	}

	return nil
}

// overwrites the secure session cookie
func revokeSessionCookie(resp http.ResponseWriter) {
	http.SetCookie(resp, &http.Cookie{
		Name:     "session", // same cookie name → overwrites session cookie
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1, // browser deletes this cookie immediatly
	})
}
