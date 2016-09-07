package main

import (
	"html/template"
	"net/http"
)

func sponsors(w http.ResponseWriter, r *http.Request) {
	log.Debug("Loading sponsors page for " + r.RemoteAddr)
	t := template.Must(template.ParseFiles(
		absPath("templates/base.html"),
		absPath("templates/sponsors.html")))
	err := t.ExecuteTemplate(w, "base", nil)
	check(err, true)
}

func sponsors_package(w http.ResponseWriter, r *http.Request) {
	log.Debug("Loading sponsors_package page for " + r.RemoteAddr)
	t := template.Must(template.ParseFiles(
		absPath("templates/base.html"),
		absPath("templates/sponsors_package.html")))
	err := t.ExecuteTemplate(w, "base", nil)
	check(err, true)
}
