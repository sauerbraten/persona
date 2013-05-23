package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"log"
	"net/http"
	"text/template"
)

var (
	// knownUsers should be a database table or something when using persona in production
	knownUsers map[string]bool

	secCookie *securecookie.SecureCookie
)

func init() {
	knownUsers = map[string]bool{}

	// initialize secure cookie storage; 16 byte ~ 128 bit AES encryption
	secCookie = securecookie.New([]byte("my very secret hash key"), []byte{0x5d, 0xa5, 0xd3, 0x90, 0xc9, 0x54, 0xa1, 0xc3, 0x70, 0x00, 0x8d, 0x6d, 0xa9, 0xd1, 0x07, 0x53})
}

func landingPage(resp http.ResponseWriter, req *http.Request) {
	log.Println("/ req. by " + req.RemoteAddr)
	template.Must(template.ParseFiles("html/landing.html")).Execute(resp, getUserEmail(req))
}

func main() {
	r := mux.NewRouter()

	// css
	r.Handle("/{file:[a-z]+\\.css}", http.FileServer(http.Dir("css")))

	// javascript
	r.Handle("/{file:[a-z]+\\.js}", http.FileServer(http.Dir("js")))

	// frontpage
	r.HandleFunc("/", landingPage)

	// login/logout
	r.HandleFunc("/signin", signIn).Methods("POST")
	r.HandleFunc("/signout", signOut).Methods("POST")

	// start listening
	log.Println("server listening on localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println(err)
	}
}
