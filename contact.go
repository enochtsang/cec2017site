package main

import (
    "html/template"
    "net/http"
)

func contact(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/contact.html")))
    err := t.ExecuteTemplate(w, "base", nil)
    check(err)
}

func contact_faq(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/faq.html")))
    err := t.ExecuteTemplate(w, "base", nil)
    check(err)
}
