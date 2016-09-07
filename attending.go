package main

import (
    "html/template"
    "net/http"
)

func attending(w http.ResponseWriter, r *http.Request) {
    log.Debug("Loading attending page for " + r.RemoteAddr)
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/attending.html")))
    err := t.ExecuteTemplate(w, "base", nil)
    check(err, true)
}
