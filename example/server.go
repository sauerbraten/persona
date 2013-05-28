package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func landingPage(resp http.ResponseWriter, req *http.Request) {
	template.Must(template.ParseFiles("html/landing.html")).Execute(resp, getEmail(req))
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
