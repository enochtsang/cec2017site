package main

import (
    "html/template"
    "net/http"
)

// var log defined in cec2017site.go

type Competition struct {
    ID    string
    Title string
    Blurb string
}

func competitions(w http.ResponseWriter, r *http.Request) {
    log.Debug("Loading competitions page for " + *r.RemoteAddr)
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/competitions.html")))
    blurb :=
        `skdjfskfjsdkldsklfjsdklfj
    skdjfskfjsdkldsklfjsdklfj
    skdjfskfjsdkldsklfjsdklfj
    skdjfskfjsdkldsklfjsdklfj
    skdjfskfjsdkldsklfjsdklfj
    skdjfskfjsdkldsklfjsdklfj`

    competitions := []Competition{
        {
            ID:    "competitions-general",
            Title: "General",
            Blurb: blurb,
        },
        {
            ID:    "senior-design",
            Title: "Senior Design",
            Blurb: blurb,
        },
        {
            ID:    "junior-design",
            Title: "Junior Design",
            Blurb: blurb,
        },
        {
            ID:    "innovative-design",
            Title: "Innovative Design",
            Blurb: blurb,
        },
        {
            ID:    "re-engineering",
            Title: "Re Engineering",
            Blurb: blurb,
        },
        {
            ID:    "engineering-communication",
            Title: "Engineering Communication",
            Blurb: blurb,
        },
        {
            ID:    "consulting",
            Title: "Consulting",
            Blurb: blurb,
        },
        {
            ID:    "extemporaneous-debate",
            Title: "Extemporaneous Debate",
            Blurb: blurb,
        },
    }
    err := t.ExecuteTemplate(w, "base", competitions)
    check(err, true)
}
