package main

import (
    "html/template"
    "log"
    "net/http"
    "os"
    "path"
    "path/filepath"
)

type Competition struct {
    ID    string
    Title string
    Blurb string
}

func check(err error) {
    if err != nil {
        panic(err)
        log.Fatal(err)
    }
}

func absPath(relativePath string) string {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    check(err)
    result := path.Join(dir, relativePath)
    log.Println("Created absPath " + result)
    return result
}

func redirectToHome(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/home", 301)
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
    err := t.ExecuteTemplate(w, "base", nil)
    check(err)
}

func home_uofc(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/uofc.html"),
        absPath("templates/widgets/didyouknow.html"),
        absPath("templates/widgets/location.html")))
    err := t.ExecuteTemplate(w, "base", nil)
    check(err)
}

func home_calgary(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/calgary.html"),
        absPath("templates/widgets/didyouknow.html"),
        absPath("templates/widgets/location.html")))
    err := t.ExecuteTemplate(w, "base", nil)
    check(err)
}

func home_hotel(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/hotel.html"),
        absPath("templates/widgets/location.html")))
    err := t.ExecuteTemplate(w, "base", nil)
    check(err)
}

func competitions(w http.ResponseWriter, r *http.Request) {
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
    check(err)
}

func sponsors(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/sponsors.html")))
    err := t.ExecuteTemplate(w, "base", nil)
    check(err)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, absPath("resources/images/favicon.ico"))
}

func main() {
    http.HandleFunc("/", redirectToHome)
    http.HandleFunc("/home", home)
    http.HandleFunc("/home/committee", home_committee)
    http.HandleFunc("/home/uofc", home_uofc)
    http.HandleFunc("/home/calgary", home_calgary)
    http.HandleFunc("/home/hotel", home_hotel)
    http.HandleFunc("/competitions", competitions)
    http.HandleFunc("/sponsors", sponsors)
    http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir(absPath("resources")))))
    http.HandleFunc("/favicon.ico", faviconHandler)

    log.Println("CEC 2017 Site Running on 8443")
    err := http.ListenAndServeTLS(":8443", absPath("temp.crt"), absPath("temp.key"), nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
