package main

import (
    "html/template"
    "net/http"
)

// CommitteeMember is info for each committee member
type CommitteeMember struct {
    Name        string
    Class       string
    Role        string
    Phone       string
    Email       string
    PicturePath string
    Info        string
}

type Location struct {
    LocationSource string
}

func home(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/home.html")))
    err := t.ExecuteTemplate(w, "base", nil)
    check(err)
}

func home_committee(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/committee.html")))
    committee := []CommitteeMember{
        {
            Name:        "Christine Cao",
            Class:       "christine-cao",
            Role:        "Chair",
            Phone:       "403-123-4567",
            Email:       "something@example.com",
            PicturePath: "/resources/images/placeholders/directors-placeholder.jpg",
            Info: ` The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.`,
        },
        {
            Name:        "Enoch Tsang",
            Class:       "enoch-tsang",
            Role:        "Executive Manager",
            Phone:       "403-123-4567",
            Email:       "something@example.com",
            PicturePath: "/resources/images/placeholders/directors-placeholder.jpg",
            Info: ` The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.`,
        },
        {
            Name:        "Byron Seto",
            Class:       "byron-seto",
            Role:        "Executive Manager",
            Phone:       "403-123-4567",
            Email:       "something@example.com",
            PicturePath: "/resources/images/placeholders/directors-placeholder.jpg",
            Info: ` The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.`,
        },
        {
            Name:        "Lami Opadokun",
            Class:       "lami-opadokun",
            Role:        "Executive Manager",
            Phone:       "403-123-4567",
            Email:       "something@example.com",
            PicturePath: "/resources/images/placeholders/directors-placeholder.jpg",
            Info: ` The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.`,
        },
        {
            Name:        "Henry Tran",
            Class:       "henry-tran",
            Role:        "Executive Manager",
            Phone:       "403-123-4567",
            Email:       "something@example.com",
            PicturePath: "/resources/images/placeholders/directors-placeholder.jpg",
            Info: ` The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.`,
        },
        {
            Name:        "Shawn Ayers",
            Class:       "shawn-ayers",
            Role:        "Executive Manager",
            Phone:       "403-123-4567",
            Email:       "something@example.com",
            PicturePath: "/resources/images/placeholders/directors-placeholder.jpg",
            Info: ` The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.`,
        },
        {
            Name:        "David Gamba",
            Class:       "david-gamba",
            Role:        "Executive Manager",
            Phone:       "403-123-4567",
            Email:       "something@example.com",
            PicturePath: "/resources/images/placeholders/directors-placeholder.jpg",
            Info: ` The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.`,
        },
        {
            Name:        "Justin Franco",
            Class:       "justin-franco",
            Role:        "Executive Manager",
            Phone:       "403-123-4567",
            Email:       "something@example.com",
            PicturePath: "/resources/images/placeholders/directors-placeholder.jpg",
            Info: ` The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.`,
        },
        {
            Name:        "Eunice Wong",
            Class:       "eunice-wong",
            Role:        "Executive Manager",
            Phone:       "403-123-4567",
            Email:       "something@example.com",
            PicturePath: "/resources/images/placeholders/directors-placeholder.jpg",
            Info: ` The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.`,
        },
        {
            Name:        "Sally Ingrid",
            Class:       "sally-ingrid",
            Role:        "Executive Manager",
            Phone:       "403-123-4567",
            Email:       "something@example.com",
            PicturePath: "/resources/images/placeholders/directors-placeholder.jpg",
            Info: ` The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.
                    The quick brown fox jumps over the lazy dog.`,
        },
    }
    err := t.ExecuteTemplate(w, "base", committee)
    check(err)
}

func home_uofc(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/uofc.html"),
        absPath("templates/widgets/didyouknow.html"),
        absPath("templates/widgets/location.html")))
    location := Location{"https://www.google.com/maps/embed/v1/place?q=univeristy%20of%20calgary&key=AIzaSyDtvcBmchfCPR_BxsOhP8UWOvamaNEQbQA&zoom=12"}
    err := t.ExecuteTemplate(w, "base", location)
    check(err)
}

func home_calgary(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/calgary.html"),
        absPath("templates/widgets/didyouknow.html"),
        absPath("templates/widgets/location.html")))
    location := Location{"https://www.google.com/maps/embed/v1/place?q=calgary%20ab&key=AIzaSyDtvcBmchfCPR_BxsOhP8UWOvamaNEQbQA"}
    err := t.ExecuteTemplate(w, "base", location)
    check(err)
}

func home_hotel(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/hotel.html"),
        absPath("templates/widgets/location.html")))
    location := Location{"https://www.google.com/maps/embed/v1/place?q=radisson%20hotel%20calgary&key=AIzaSyDtvcBmchfCPR_BxsOhP8UWOvamaNEQbQA"}
    err := t.ExecuteTemplate(w, "base", location)
    check(err)
}
