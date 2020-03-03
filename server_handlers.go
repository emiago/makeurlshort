package main

import (
	"flixwebtest/shorturl"
	"flixwebtest/urlshorty"
	"fmt"
	"net/http"
	"net/url"
	"text/template"
)

type UrlInputDetails struct {
	Error     string
	Urlinput  string
	ShortUrls []shorturl.UrlData
}

func SaveShortUrl(surl string, original string) error {
	s := shorturl.UrlData{
		OriginalUrl: original,
		Url:         surl,
	}

	return shorturl.PostShortUrl(s)
}

func ParseInputUrl(inputurl string) error {
	shurl, err := urlshorty.Parse(inputurl)
	if err != nil {
		return err
	}

	return SaveShortUrl(shurl, inputurl)
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(layout.ParseFiles("templates/form.html"))

	shortUrls, _ := shorturl.GetShortUrls()
	details := UrlInputDetails{
		ShortUrls: shortUrls, //I need to load this every time
	}

	if r.Method != http.MethodPost {
		tmpl.Execute(w, details)
		return
	}

	inputurl := r.FormValue("urlinput")

	if err := ParseInputUrl(inputurl); err != nil {
		details.Error = fmt.Sprintf("Parsing url failed. %s", err.Error())
		tmpl.Execute(w, details)
		return
	}

	// shurl, err := urlshorty.Parse(inputurl)
	// if err != nil {
	// 	details.Error = fmt.Sprintf("Parsing url failed. %s", err.Error())
	// 	tmpl.Execute(w, details)
	// 	return
	// }

	// SaveShortUrl(shurl, inputurl)
	shortUrls, _ = shorturl.GetShortUrls()

	details = UrlInputDetails{
		Urlinput:  inputurl,
		ShortUrls: shortUrls,
	}

	tmpl.Execute(w, details)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	ruri := r.URL.RequestURI()

	shortUrls, _ := shorturl.GetShortUrls() //This should be refactored by using map and key as url
	for _, s := range shortUrls {
		myurl, _ := url.Parse(s.Url)
		if myurl.RequestURI() == ruri {
			http.Redirect(w, r, s.OriginalUrl, 301)
			return
		}
	}

	http.Error(w, "not found", http.StatusNotFound)
}
