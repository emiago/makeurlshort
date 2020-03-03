package main

import (
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	layout *template.Template
)

func corsHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == "OPTIONS" {
			//handle preflight in here
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		} else {
			h(w, r)
		}
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == "OPTIONS" {
			//handle preflight in here
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func serverRun() {
	// r := http.NewServeMux()
	r := mux.NewRouter()

	layout = template.Must(template.ParseFiles("templates/layout.html"))

	r.HandleFunc("/", FormHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// http.ListenAndServe(":8080", r)
	log.Fatal(srv.ListenAndServe())
}
