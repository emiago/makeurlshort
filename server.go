package main

import (
	"encoding/json"
	"flixwebtest/urlshorty"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	layout *template.Template
)

type UrlInputDetails struct {
	Error     string
	Submited  bool
	Urlinput  string
	ShortUrls []ShortUrl
}

type JsonMessage struct {
	Id      int
	Name    string
	Message string
}

// type messageHandler struct {
// 	message string
// }

// func (m *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, m.message)
// }

func JSONHandler(data interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Failed to parse message", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(b)
	}
}

var (
	shortUrls []ShortUrl
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(layout.ParseFiles("templates/form.html"))

	details := UrlInputDetails{
		Submited:  false,
		Urlinput:  "",
		ShortUrls: shortUrls, //I need to load this every time
	}

	if r.Method != http.MethodPost {
		tmpl.Execute(w, details)
		return
	}

	inputurl := r.FormValue("urlinput")

	shorturl, err := urlshorty.Parse(inputurl)
	if err != nil {
		details.Error = fmt.Sprintf("Parsing url failed. %s", err.Error())
		tmpl.Execute(w, details)
		return
	}

	surl := ShortUrl{
		OriginalUrl: inputurl,
		Url:         shorturl,
	}
	shortUrls = append(shortUrls, surl)

	details = UrlInputDetails{
		Submited: true,
		Urlinput: inputurl,
	}

	details.ShortUrls = append(details.ShortUrls, shortUrls...)

	tmpl.Execute(w, details)
}

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
	// fs := http.FileServer(http.Dir("./static"))
	// r.Handle("/static/", http.StripPrefix("/static/", fs))

	// r.HandleFunc("/", HomeHandler)
	// r.HandleFunc("/home", HomeHandler)
	r.HandleFunc("/", FormHandler)

	// r.Handle("/api", &messageHandler{"Welcome to API"})
	// api := r.PathPrefix("/api").Subrouter()
	// api.Use(corsMiddleware)
	// api.HandleFunc("/json", JSONHandler(JsonMessage{
	// 	Id:      1,
	// 	Name:    "M1",
	// 	Message: "Some json with message",
	// }))

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
